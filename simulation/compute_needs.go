package simulation

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/game"
)

type Needs struct {
	ID             int16
	AvailableUnits int16
	SupportDemands []SupportDemand
}

// Needed units: units needed to avoid the loss of the planet.
//
// Left: the number of turns left before the planet is lost.
//
// PenaltyAfter: the penalty of units needed for each turn after the loss of the planet (i.e. growth).
//
// UselessBefore: tells if it is useless to send fleets before <Left> turns. Useful when the enemy sends more than MaxUnits,
// because in that case the planet will be lost anyway even if Units = MaxUnits.
//
type SupportDemand struct {
	NeededUnits   int16
	Left          int16
	PenaltyAfter  int16
	UselessBefore bool
}

func ComputeNeeds(gameMap *game.Map, planetID int16) Needs {
	timeline := engine.CreateTimelineForPlanets(gameMap, []int16{planetID})

	if timeline.PlanetTimelinesMap[planetID] == nil {
		return Needs{
			ID:             0,
			AvailableUnits: 0,
			SupportDemands: make([]SupportDemand, 0),
		}
	}

	var lostUnits int16
	maxAvailableUnits := timeline.PlanetTimelinesMap[planetID].CurrentTurn().Units
	maxUnits := timeline.PlanetTimelinesMap[planetID].CurrentTurn().MaxUnits
	minUnits := maxUnits
	hasBeenAttacked := false
	supportDemands := make([]SupportDemand, 0)

	for i := 0; i < common.SIMULATION_HORIZON; i++ {
		timeline.NextTurn()

		currentTurn := timeline.PlanetTimelinesMap[planetID].CurrentTurn()
		previousTurn := timeline.PlanetTimelinesMap[planetID].PreviousTurn()

		if maxUnits-previousTurn.Units < currentTurn.Growth && !hasBeenAttacked {
			lostUnits += currentTurn.Growth - (maxUnits - previousTurn.Units)
		}

		if currentTurn.OwnerID != previousTurn.OwnerID {
			supportDemands = append(supportDemands, SupportDemand{
				NeededUnits:   currentTurn.Units + 1,
				Left:          int16(timeline.Turn),
				UselessBefore: currentTurn.Units+previousTurn.Units >= currentTurn.MaxUnits,
				PenaltyAfter:  currentTurn.Growth,
			})

			// Simulate the arrival of a support fleet to resume the simulation
			// The support fleet leaves the planet at 1 unit by default
			turnCopy := currentTurn.Copy()
			turnCopy.OwnerID = previousTurn.OwnerID
			turnCopy.Units = 1

			timeline.PlanetTimelinesMap[planetID].SetCurrentTurn(turnCopy)

			// The planet is lost after some turns, it cannot have units
			minUnits = 0
			hasBeenAttacked = true

		} else if currentTurn.Units < previousTurn.Units && currentTurn.Units < minUnits {
			minUnits = currentTurn.Units
			hasBeenAttacked = true
		}
	}

	availableUnits := minUnits + lostUnits

	if availableUnits > maxAvailableUnits {
		availableUnits = maxAvailableUnits
	}

	if availableUnits > 0 {
		availableUnits -= 1
	}

	return Needs{
		ID:             planetID,
		AvailableUnits: availableUnits,
		SupportDemands: supportDemands,
	}
}

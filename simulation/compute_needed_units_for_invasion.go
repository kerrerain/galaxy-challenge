package simulation

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/game"
)

func ComputeNeededUnitsForInvasion(gameMap *game.Map, sourceID int16, targetID int16) int16 {
	timeline := engine.CreateTimelineForPlanets(gameMap, []int16{targetID})
	distanceInTurns := gameMap.DistanceMap[sourceID][targetID].Turns

	for i := 0; i < int(distanceInTurns); i++ {
		timeline.NextTurn()
	}

	unitsNeeded := timeline.PlanetTimelinesMap[targetID].CurrentTurn().Units + 1

	if unitsNeeded < common.MIN_FLEET_UNITS {
		unitsNeeded = common.MIN_FLEET_UNITS
	}

	return unitsNeeded
}

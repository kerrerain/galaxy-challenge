package engine

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
)

type PlanetTimeline struct {
	ID    uint16
	Turns []dto.StatusPlanet
}

func CreatePlanetTimeline(planet dto.StatusPlanet) *PlanetTimeline {
	return &PlanetTimeline{
		ID:    planet.ID,
		Turns: []dto.StatusPlanet{planet},
	}
}

func CreatePlanetTimelines(planets []dto.StatusPlanet) []*PlanetTimeline {
	planetTimelines := make([]*PlanetTimeline, len(planets))

	for index, planet := range planets {
		planetTimelines[index] = CreatePlanetTimeline(planet)
	}

	return planetTimelines
}

func (p *PlanetTimeline) NextTurn(fleets []dto.StatusFleet) {
	planet := p.Turns[len(p.Turns)-1]
	nextTurnPlanet := planet.Copy()

	applyGrowth(&nextTurnPlanet)
	applyFleets(&nextTurnPlanet, fleets)

	p.Turns = append(p.Turns, nextTurnPlanet)
}

// The ownership tells how much a planet is owned by the player at the end of the turns.
// Negative values means that an enemy owns the planet.
func (p PlanetTimeline) Ownership() int {
	if len(p.Turns) == 0 {
		return 0
	}

	lastTurn := p.Turns[len(p.Turns)-1]

	if lastTurn.OwnerID != common.PLAYER_OWNER_ID {
		return -1 * int(lastTurn.Units)
	} else {
		return int(lastTurn.Units)
	}
}

func applyGrowth(planet *dto.StatusPlanet) {
	if planet.OwnerID != common.NEUTRAL_OWNER_ID {
		planet.Units = common.Min(planet.MaxUnits, planet.Units+planet.Growth)
	}
}

func applyFleets(planet *dto.StatusPlanet, fleets []dto.StatusFleet) {
	for _, fleet := range fleets {
		planet.Units = planet.Units - fleet.Units

		if planet.Units < 0 {
			planet.OwnerID = fleet.OwnerID // The other player has earned the planet
			planet.Units = -1 * planet.Units
		}
	}
}

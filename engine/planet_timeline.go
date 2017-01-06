package engine

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
)

type PlanetTimeline struct {
	Turns []dto.StatusPlanet
}

func (p *PlanetTimeline) NextTurn(fleets []dto.StatusFleet) {
	planet := p.Turns[len(p.Turns)-1]
	nextTurnPlanet := planet.Copy()

	applyGrowth(&nextTurnPlanet)
	applyFleets(&nextTurnPlanet, fleets)

	p.Turns = append(p.Turns, nextTurnPlanet)
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

func CreatePlanetTimeline(planet dto.StatusPlanet) *PlanetTimeline {
	return &PlanetTimeline{
		Turns: []dto.StatusPlanet{planet},
	}
}

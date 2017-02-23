package engine

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
)

type PlanetTimeline struct {
	ID                  int16
	Turns               []dto.StatusPlanet
	Origin              *dto.StatusPlanet
	TotalUnitsSent      int16
	TotalEnemyUnitsSent int16
}

func CreatePlanetTimeline(planet dto.StatusPlanet) *PlanetTimeline {
	return &PlanetTimeline{
		ID:     planet.ID,
		Turns:  []dto.StatusPlanet{planet},
		Origin: &planet,
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
	applyFleetsToTotal(p, fleets)

	p.Turns = append(p.Turns, nextTurnPlanet)
}

func (p PlanetTimeline) CurrentTurn() dto.StatusPlanet {
	return p.Turns[len(p.Turns)-1]
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

func applyFleetsToTotal(planetTimeline *PlanetTimeline, fleets []dto.StatusFleet) {
	for _, fleet := range fleets {
		if fleet.OwnerID == common.PLAYER_OWNER_ID {
			planetTimeline.TotalUnitsSent += fleet.Units
		} else {
			planetTimeline.TotalEnemyUnitsSent += fleet.Units
		}
	}
}

package engine

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"log"
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

func (p PlanetTimeline) SetCurrentTurn(planet dto.StatusPlanet) {
	p.Turns[len(p.Turns)-1] = planet
}

func (p PlanetTimeline) PreviousTurn() dto.StatusPlanet {
	return p.Turns[len(p.Turns)-2]
}

func applyGrowth(planet *dto.StatusPlanet) {
	formerUnits := planet.Units

	if planet.OwnerID != common.NEUTRAL_OWNER_ID {
		planet.Units = common.Min(planet.MaxUnits, formerUnits+planet.Growth)
	}

	if common.DEBUG_MODE && planet.Units != formerUnits {
		log.Printf("Growth on planet %d: from %d to %d", planet.ID, formerUnits, planet.Units)
	}
}

func applyFleets(planet *dto.StatusPlanet, fleets []dto.StatusFleet) {
	for _, fleet := range fleets {

		if common.DEBUG_MODE {
			log.Printf("Fleet arrival on planet %d: player %d sent %d units", planet.ID, fleet.OwnerID, fleet.Units)
		}

		if planet.OwnerID == fleet.OwnerID {
			planet.Units += fleet.Units
		} else {
			planet.Units -= fleet.Units
		}

		previousOwnerID := planet.OwnerID

		if planet.Units < 0 {
			planet.OwnerID = fleet.OwnerID // The other player has earned the planet
			planet.Units = -1 * planet.Units

			if common.DEBUG_MODE {
				log.Printf("Ownership has changed: from player %d to %d, %d units", previousOwnerID, fleet.OwnerID, planet.Units)
			}
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

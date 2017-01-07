package engine

import (
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
)

type Timeline struct {
	PlanetTimelines []*PlanetTimeline
	FleetScheduler  *FleetScheduler
	Turn            int
}

func (t *Timeline) NextTurn() {
	t.Turn = t.Turn + 1

	for _, planet := range t.PlanetTimelines {
		fleets := t.FleetScheduler.TurnFleetsForPlanet(t.Turn, planet.ID)
		planet.NextTurn(fleets)
	}
}

func CreateTimeline(gameMap *game.Map) Timeline {
	return Timeline{
		PlanetTimelines: initPlanetTimelines(gameMap.Planets),
		FleetScheduler:  CreateFleetScheduler(gameMap.Fleets),
	}
}

func initPlanetTimelines(planets []dto.StatusPlanet) []*PlanetTimeline {
	planetTimelines := make([]*PlanetTimeline, len(planets))

	for index, planet := range planets {
		planetTimelines[index] = CreatePlanetTimeline(planet)
	}

	return planetTimelines
}

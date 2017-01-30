package engine

import (
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
)

type Timeline struct {
	GameMap         *game.Map
	PlanetTimelines []*PlanetTimeline
	FleetScheduler  *FleetScheduler
	Turn            int
}

func CreateTimeline(gameMap *game.Map) Timeline {
	return Timeline{
		GameMap:         gameMap,
		PlanetTimelines: CreatePlanetTimelines(gameMap.Planets),
		FleetScheduler:  CreateFleetScheduler(gameMap.Fleets),
	}
}

func (t *Timeline) NextTurn() {
	t.Turn = t.Turn + 1

	for _, planet := range t.PlanetTimelines {
		fleets := t.FleetScheduler.TurnFleetsForPlanet(t.Turn, planet.ID)
		planet.NextTurn(fleets)
	}
}

func (t *Timeline) ScheduleMoveForNextTurn(playerID uint16, move dto.Move) {
	for _, fleet := range move.Fleets {
		t.FleetScheduler.AddFleet(t.GameMap.MapMoveFleet(playerID, fleet))
	}
}

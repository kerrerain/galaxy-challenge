package engine

import (
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
)

type Timeline struct {
	GameMap            *game.Map
	PlanetTimelines    []*PlanetTimeline
	PlanetTimelinesMap map[int16]*PlanetTimeline
	FleetScheduler     *FleetScheduler
	Turn               int
}

func CreateTimeline(gameMap *game.Map) Timeline {
	return doCreateTimeline(gameMap, gameMap.Planets, gameMap.Fleets)
}

func CreateTimelineForPlanets(gameMap *game.Map, ids []int16) Timeline {
	planets := dto.FilterStatusPlanets(gameMap.Planets, func(planet dto.StatusPlanet) bool {
		found := false
		for _, id := range ids {
			if planet.ID == id {
				found = true
			}
		}
		return found
	})

	fleets := dto.FilterStatusFleets(gameMap.Fleets, func(fleet dto.StatusFleet) bool {
		found := false
		for _, id := range ids {
			if fleet.TargetID == id {
				found = true
			}
		}
		return found
	})

	return doCreateTimeline(gameMap, planets, fleets)
}

func doCreateTimeline(gameMap *game.Map, planets []dto.StatusPlanet, fleets []dto.StatusFleet) Timeline {
	timeline := Timeline{
		GameMap:            gameMap,
		PlanetTimelines:    CreatePlanetTimelines(planets),
		PlanetTimelinesMap: make(map[int16]*PlanetTimeline),
		FleetScheduler:     CreateFleetScheduler(fleets),
	}

	for _, planetTimeline := range timeline.PlanetTimelines {
		timeline.PlanetTimelinesMap[planetTimeline.ID] = planetTimeline
	}

	return timeline
}

func (t *Timeline) NextTurn() {
	t.Turn = t.Turn + 1

	for _, planet := range t.PlanetTimelines {
		fleets := t.FleetScheduler.TurnFleetsForPlanet(t.Turn, planet.ID)
		planet.NextTurn(fleets)
	}
}

func (t *Timeline) ScheduleMoveForNextTurn(playerID int16, move dto.Move) {
	for _, fleet := range move.Fleets {
		t.FleetScheduler.AddFleet(t.GameMap.MapMoveFleet(playerID, fleet))
	}
}

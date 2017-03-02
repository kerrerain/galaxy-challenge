package engine

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"log"
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

	if common.DEBUG_MODE {
		log.Printf("Moving to next turn: %d", t.Turn)
	}

	for _, planet := range t.PlanetTimelines {
		fleets := t.FleetScheduler.TurnFleetsForPlanet(t.Turn, planet.ID)
		planet.NextTurn(fleets)
	}
}

func (t *Timeline) ScheduleMoveForNextTurn(playerID int16, move dto.Move) {
	for _, fleet := range move.Fleets {
		// Remove units from the source of the fleet
		if t.PlanetTimelinesMap[fleet.SourceID] != nil {
			planet := t.PlanetTimelinesMap[fleet.SourceID].CurrentTurn().Copy()
			planet.Units -= fleet.Units

			t.PlanetTimelinesMap[fleet.SourceID].Turns[t.Turn] = planet

			if common.DEBUG_MODE {
				log.Printf("Sending %d units from planet %d, %d units remaining", fleet.Units, fleet.SourceID,
					planet.Units)
			}
		}

		t.FleetScheduler.AddFleet(t.GameMap.MapMoveFleet(playerID, fleet))
	}
}

func (t Timeline) Status() dto.Status {
	return dto.Status{
		Fleets:  t.FleetScheduler.Fleets(),
		Planets: t.Planets(),
	}
}

func (t Timeline) Planets() []dto.StatusPlanet {
	planets := make([]dto.StatusPlanet, len(t.PlanetTimelines))

	for i, planetTimeline := range t.PlanetTimelines {
		planets[i] = planetTimeline.CurrentTurn()
	}

	return planets
}

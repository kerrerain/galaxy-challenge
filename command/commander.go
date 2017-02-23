package command

import (
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/game"
)

type Commander struct {
	AvailableUnitsOnPlanet map[int16]int16
	Fleets                 []*dto.MoveFleet
}

func CreateCommander(gameMap *game.Map) *Commander {
	return &Commander{
		AvailableUnitsOnPlanet: initAvailableUnitsOnPlanet(gameMap),
		Fleets:                 make([]*dto.MoveFleet, 0),
	}
}

func initAvailableUnitsOnPlanet(gameMap *game.Map) map[int16]int16 {
	unitsMap := make(map[int16]int16)

	for _, planet := range gameMap.Planets {
		unitsMap[planet.ID] = planet.Units
	}

	return unitsMap
}

func CreateCommanderFromTimeline(timeline engine.Timeline) *Commander {
	return &Commander{
		AvailableUnitsOnPlanet: initAvailableUnitsOnPlanetFromTimeline(timeline),
		Fleets:                 make([]*dto.MoveFleet, 0),
	}
}

func initAvailableUnitsOnPlanetFromTimeline(timeline engine.Timeline) map[int16]int16 {
	unitsMap := make(map[int16]int16)

	for _, planet := range timeline.PlanetTimelines {
		unitsMap[planet.ID] = planet.Turns[len(planet.Turns)-1].Units
	}

	return unitsMap
}

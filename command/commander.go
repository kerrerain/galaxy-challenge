package command

import (
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
)

type Commander struct {
	AvailableUnitsOnPlanet map[uint16]int16
	Fleets                 []*dto.MoveFleet
}

func CreateCommander(gameMap *game.Map) *Commander {
	return &Commander{
		AvailableUnitsOnPlanet: initAvailableUnitsOnPlanet(gameMap),
		Fleets:                 make([]*dto.MoveFleet, 0),
	}
}

func initAvailableUnitsOnPlanet(gameMap *game.Map) map[uint16]int16 {
	unitsMap := make(map[uint16]int16)

	for _, planet := range gameMap.Planets {
		unitsMap[planet.ID] = planet.Units
	}

	return unitsMap
}

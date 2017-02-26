package amon

import (
	"github.com/magleff/galaxy-challenge/command"
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"github.com/magleff/galaxy-challenge/simulation"
	"log"
)

// Spirit of politeness and austerity.
//
func Run(gameMap *game.Map) dto.Move {
	commander := command.CreateCommander(gameMap)
	ownPlanets := dto.FilterOwnPlanets(gameMap.Planets)

	for _, sourcePlanet := range ownPlanets {
		nearestPlanets := gameMap.NearestPlanetsMap[sourcePlanet.ID]
		targetPlanet := chooseTarget(gameMap, nearestPlanets)

		log.Printf("Planet %d targets %d", sourcePlanet.ID, targetPlanet.ID)

		if targetPlanet.ID != 0 {
			order := command.Order{
				SourceID: sourcePlanet.ID,
				TargetID: targetPlanet.ID,
				Units:    simulation.ComputeNeededUnitsForInvasion(gameMap, sourcePlanet.ID, targetPlanet.ID),
			}

			if simulation.ComputeKept(gameMap, sourcePlanet.ID, order) {
				commander.SendOrder(order)
			} else {
				log.Printf("Keeping units on planet %d to avoid losing it", sourcePlanet.ID)
			}
		} else {
			log.Printf(`Keeping units on planet %d because there are no
				more targets on the map`, sourcePlanet.ID)
		}
	}

	return commander.BuildMove()
}

func chooseTarget(gameMap *game.Map, nearestPlanets []int16) dto.StatusPlanet {
	target := dto.StatusPlanet{
		ID:      0,
		OwnerID: 0,
	}

	for _, targetID := range nearestPlanets {
		if target.ID == 0 &&
			target.OwnerID != common.PLAYER_OWNER_ID &&
			!simulation.ComputeTaken(gameMap, targetID) {

			target = dto.GetByID(gameMap.Planets, targetID)
		}
	}

	return target
}

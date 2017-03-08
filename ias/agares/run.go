package agares

import (
	"github.com/magleff/galaxy-challenge/command"
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"github.com/magleff/galaxy-challenge/simulation"
	"log"
)

const NAME = "AGARES"

// Spirit of Fortitude and Creation.
//
// The evolution of Amon.
//
func Run(gameMap *game.Map, playerID int16) dto.Move {
	commander := command.CreateCommander(gameMap)
	ownPlanets := dto.FilterPlanetsByPlayerID(gameMap.Planets, playerID)

	for _, sourcePlanet := range ownPlanets {
		needs := simulation.ComputeNeeds(gameMap, sourcePlanet.ID)
		log.Printf("Needs for planet %d: %v", sourcePlanet.ID, needs)

		nearestPlanets := gameMap.NearestPlanetsMap[sourcePlanet.ID]
		targetPlanet := chooseTarget(gameMap, nearestPlanets, playerID)

		log.Printf("Planet %d targets %d", sourcePlanet.ID, targetPlanet.ID)

		if targetPlanet.ID != 0 {
			unitsSent := simulation.ComputeNeededUnitsForInvasion(gameMap, sourcePlanet.ID, targetPlanet.ID)

			if unitsSent > needs.AvailableUnits {
				unitsSent = needs.AvailableUnits
			}

			order := command.Order{
				SourceID: sourcePlanet.ID,
				TargetID: targetPlanet.ID,
				Units:    unitsSent,
			}
			commander.SendOrder(order)
		}
	}

	return commander.BuildMove()
}

func chooseTarget(gameMap *game.Map, nearestPlanets []int16, playerID int16) dto.StatusPlanet {
	target := dto.StatusPlanet{
		ID: 0,
	}

	for _, targetID := range nearestPlanets {
		if target.ID == 0 && !simulation.ComputeTaken(gameMap, targetID, playerID) {
			planet := dto.GetByID(gameMap.Planets, targetID)
			if planet.OwnerID != playerID &&
				(planet.Growth > 0 || planet.OwnerID != common.NEUTRAL_OWNER_ID) {

				target = planet
			}
		}
	}

	return target
}

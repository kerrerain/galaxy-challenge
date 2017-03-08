package vepar

import (
	"github.com/magleff/galaxy-challenge/command"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"github.com/magleff/galaxy-challenge/simulation"
	"log"
)

const NAME = "VEPAR"

// Spirit of Liberty and Cunning.
//
// The evolution of Agares.
//
func Run(gameMap *game.Map, playerID int16) dto.Move {
	commander := command.CreateCommander(gameMap)
	ownPlanets := dto.FilterPlanetsByPlayerID(gameMap.Planets, playerID)

	for _, sourcePlanet := range ownPlanets {
		var targetPlanet dto.StatusPlanet

		needs := simulation.ComputeNeeds(gameMap, sourcePlanet.ID)
		log.Printf("Needs for planet %d: %v", sourcePlanet.ID, needs)

		nearestPlanets := gameMap.NearestPlanetsMap[sourcePlanet.ID]
		targetPlanet = chooseROITarget(gameMap, sourcePlanet, nearestPlanets, playerID)

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

func chooseROITarget(gameMap *game.Map, sourcePlanet dto.StatusPlanet, nearestPlanets []int16, playerID int16) dto.StatusPlanet {
	target := dto.StatusPlanet{
		ID: 0,
	}

	minRoi := 0.0

	for _, targetID := range nearestPlanets {
		roi := simulation.ComputeROI(gameMap, sourcePlanet.ID, targetID, playerID)
		taken := simulation.ComputeTaken(gameMap, targetID, playerID)

		if roi != 0 && (roi < minRoi || minRoi == 0) && !taken {
			minRoi = roi
			target = dto.GetByID(gameMap.Planets, targetID)
		}
	}

	return target
}

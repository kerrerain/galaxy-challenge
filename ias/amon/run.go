package amon

import (
	"github.com/magleff/galaxy-challenge/command"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"github.com/magleff/galaxy-challenge/simulation"
	"log"
)

func Run(gameMap *game.Map) dto.Move {
	commander := command.CreateCommander(gameMap)
	ownPlanets := dto.FilterOwnPlanets(gameMap.Planets)

	for _, sourcePlanet := range ownPlanets {
		nearestPlanets := gameMap.NearestPlanetsMap[sourcePlanet.ID]
		targetPlanet := chooseTarget(gameMap, nearestPlanets)

		commander.SendOrder(command.Order{
			SourceID: sourcePlanet.ID,
			TargetID: targetPlanet.ID,
			Units:    int16(targetPlanet.Units) + 1,
		})

		log.Printf("Nearest planet to take from %d: %d", sourcePlanet.ID, targetPlanet.ID)
	}

	return commander.BuildMove()
}

func chooseTarget(gameMap *game.Map, nearestPlanets []int16) dto.StatusPlanet {
	target := dto.StatusPlanet{
		ID: 0,
	}

	for _, targetID := range nearestPlanets {
		if target.ID == 0 && !simulation.ComputeTaken(gameMap, targetID) {
			target = dto.GetByID(gameMap.Planets, targetID)
		}
	}

	return target
}

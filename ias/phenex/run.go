package phenex

import (
	"github.com/magleff/galaxy-challenge/command"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/evaluation"
	"github.com/magleff/galaxy-challenge/game"
	"github.com/magleff/galaxy-challenge/simulation"
	"log"
	"sort"
)

// Spirit of kindness and mediation.
//
// Selects planets from the front to attack,
// and planets from the rear to support the front.
//
// Simulate the future of potential targets,
// and avoids losing a planet at all costs.
//
// Better used in the middle of a game,
// In order to take efficiently enemy planets.
//
func Run(gameMap *game.Map) dto.Move {
	timeline := engine.CreateTimeline(gameMap)
	result := evaluation.Run(timeline, evaluation.Configuration{
		NumberFrontPlanets: 5,
	})

	commander := command.CreateCommander(gameMap)
	enemyIDs := enemyPlanetIDs(gameMap.Planets)
	sourceIDs := sourcePlanetIDs(result.FrontPlanets)
	targetIDs := enemyIDs

	log.Printf("enemyIDs %v", enemyIDs)

	if len(enemyIDs) > 0 {
		enemyPlanetsInvasionCost := simulation.ComputeFullInvasionCost(gameMap, targetIDs, sourceIDs, enemyIDs)
		sort.Sort(simulation.ByLowestInvasionCost(enemyPlanetsInvasionCost))
		target := enemyPlanetsInvasionCost[0]

		log.Printf("Costs %v", enemyPlanetsInvasionCost)
		log.Printf("All out attack on planet %d", target.ID)

		commander.AllOutAttackOnPlanet(target.ID, sourceIDs)
	}
	return commander.BuildMove()
}

func enemyPlanetIDs(planets []dto.StatusPlanet) []int16 {
	enemyPlanets := dto.FilterEnemyPlanets(planets)
	ids := make([]int16, len(enemyPlanets))

	for i, planet := range enemyPlanets {
		ids[i] = planet.ID
	}

	return ids
}

func sourcePlanetIDs(planets []*evaluation.ResultPlanet) []int16 {
	ids := make([]int16, len(planets))

	for i, planet := range planets {
		ids[i] = planet.ID
	}

	return ids
}

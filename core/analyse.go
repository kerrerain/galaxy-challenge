package core

import (
	"github.com/magleff/galaxy-challenge/models"
	"log"
)

func Analyse(universe models.Universe) (models.Response, error) {
	response := models.CreateResponse()

	// 1 -> Target planets to conquer
	targets := targetPlanets(universe)

	log.Println(len(targets), " possible planets to target")

	// 2 -> Sort them by worth

	// 3 -> Build fleets
	ownPlanets := filterOwnPlanets(planets)
	fleets := buildFleets(targets, ownPlanets)

	log.Println(len(fleets), " engaged in battle")

	// 4 -> Trigger terraformations

	return response, nil
}

func filterOwnPlanets(planets []Planet) []Planet {
	return models.FilterPlanets(planets, func(planet Planet) bool {
		return planet.OwnerID == PLAYER_OWNER_ID
	})
}

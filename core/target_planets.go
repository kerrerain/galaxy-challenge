package core

import (
	"github.com/magleff/galaxy-challenge/models"
)

func targetPlanets(universe models.Universe) []models.Planet {
	toFilter := universe.Planets
	ownFleets := universe.Fleets

	filtered := make([]models.Planet, 0)

	for _, planet := range toFilter {
		if targetPlanetsPredicate(planet, ownFleets) {
			filtered = append(filtered, planet)
		}
	}

	return filtered
}

func targetPlanetsPredicate(planet models.Planet, ownFleets []models.Fleet) bool {
	return planet.OwnerID != PLAYER_OWNER_ID &&
		planet.Habitable() &&
		!isPlanetTargetedByOwnFleet(planet, ownFleets)
}

func isPlanetTargetedByOwnFleet(planet models.Planet, ownFleets []models.Fleet) bool {
	return false
}

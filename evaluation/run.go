package evaluation

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/engine"
)

func Run(timeline engine.Timeline, config Configuration) Result {
	planets := createResultPlanets(timeline)
	ownPlanets := FilterResultPlanets(planets, func(planet *ResultPlanet) bool {
		return planet.Origin.OwnerID == common.PLAYER_OWNER_ID
	})

	result := &Result{
		Planets:        planets,
		FrontLimit:     ComputeFrontLimit(ownPlanets, config.NumberFrontPlanets),
		FrontPlanets:   make([]*ResultPlanet, 0),
		SupportPlanets: make([]*ResultPlanet, 0),
	}

	frontAndSupportPlanets(result)

	return *result
}

func frontAndSupportPlanets(result *Result) {
	for _, planet := range result.Planets {
		if planet.Origin.OwnerID == common.PLAYER_OWNER_ID {
			if planet.DistanceToEnemy <= result.FrontLimit {
				result.FrontPlanets = append(result.FrontPlanets, planet)
			} else {
				result.SupportPlanets = append(result.SupportPlanets, planet)
			}
		}
	}
}

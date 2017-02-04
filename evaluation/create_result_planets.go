package evaluation

import (
	"github.com/magleff/galaxy-challenge/engine"
)

func createResultPlanets(timeline engine.Timeline) []*ResultPlanet {
	resultPlanets := make([]*ResultPlanet, len(timeline.PlanetTimelines))

	for i, planet := range timeline.PlanetTimelines {
		resultPlanets[i] = &ResultPlanet{
			Origin:          planet.Origin,
			Loss:            ComputeLoss(planet),
			DistanceToEnemy: timeline.GameMap.ComputeMedianDistance(planet.Origin.ID).Raw,
		}
	}

	return resultPlanets
}

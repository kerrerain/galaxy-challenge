package evaluation

import (
	"sort"
)

func ComputeFrontLimit(planets []*ResultPlanet, maxFrontPlanets int) uint16 {
	sort.Sort(ByLowestDistanceToEnemy(planets))

	length := len(planets)

	if length == 0 {
		return 0
	} else if maxFrontPlanets >= length {
		return planets[length-1].DistanceToEnemy
	} else {
		return planets[maxFrontPlanets-1].DistanceToEnemy
	}
}

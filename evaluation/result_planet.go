package evaluation

import (
	"github.com/magleff/galaxy-challenge/dto"
)

type ResultPlanet struct {
	ID               int16
	Origin           *dto.StatusPlanet
	Loss             int
	Growth           int16
	DistanceToEnemy  int16
	DistanceToPlayer int16
}

type ByLowestDistanceToEnemy []*ResultPlanet

func (a ByLowestDistanceToEnemy) Len() int      { return len(a) }
func (a ByLowestDistanceToEnemy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByLowestDistanceToEnemy) Less(i, j int) bool {
	return a[i].DistanceToEnemy < a[j].DistanceToEnemy
}

type ByLowestDistanceToPlayer []*ResultPlanet

func (a ByLowestDistanceToPlayer) Len() int      { return len(a) }
func (a ByLowestDistanceToPlayer) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByLowestDistanceToPlayer) Less(i, j int) bool {
	return a[i].DistanceToPlayer < a[j].DistanceToPlayer
}

func FilterResultPlanets(toFilter []*ResultPlanet, predicate func(planet *ResultPlanet) bool) []*ResultPlanet {
	filtered := make([]*ResultPlanet, 0)
	for _, planet := range toFilter {
		if predicate(planet) {
			filtered = append(filtered, planet)
		}
	}
	return filtered
}

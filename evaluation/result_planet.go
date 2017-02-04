package evaluation

import (
	"github.com/magleff/galaxy-challenge/dto"
)

type ResultPlanet struct {
	Loss            int
	DistanceToEnemy float64
	Origin          *dto.StatusPlanet
}

type ByLowestDistanceToEnemy []*ResultPlanet

func (a ByLowestDistanceToEnemy) Len() int      { return len(a) }
func (a ByLowestDistanceToEnemy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByLowestDistanceToEnemy) Less(i, j int) bool {
	return a[i].DistanceToEnemy < a[j].DistanceToEnemy
}

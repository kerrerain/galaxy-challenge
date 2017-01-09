package paimon

import (
	"github.com/magleff/galaxy-challenge/dto"
)

type PlanetEvaluation struct {
	Score  int
	Planet dto.StatusPlanet
}

type ByGreatestScore []*PlanetEvaluation

func (a ByGreatestScore) Len() int           { return len(a) }
func (a ByGreatestScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByGreatestScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

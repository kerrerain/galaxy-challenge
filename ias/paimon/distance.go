package paimon

import (
	"github.com/magleff/galaxy-challenge/dto"
)

type Distance struct {
	Distance float64
	Planet   dto.StatusPlanet
}

type ByLowestDistance []Distance

func (a ByLowestDistance) Len() int           { return len(a) }
func (a ByLowestDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLowestDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }

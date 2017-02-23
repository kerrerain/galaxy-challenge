package game

import (
	"github.com/magleff/galaxy-challenge/dto"
)

type Map struct {
	Planets           []dto.StatusPlanet
	Fleets            []dto.StatusFleet
	DistanceMap       map[int16]map[int16]Distance
	NearestPlanetsMap map[int16][]int16
	Turn              int16
	Initialized       bool
}

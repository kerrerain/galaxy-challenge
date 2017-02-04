package game

import (
	"github.com/magleff/galaxy-challenge/dto"
)

type Map struct {
	Planets     []dto.StatusPlanet
	Fleets      []dto.StatusFleet
	DistanceMap map[uint16]map[uint16]Distance
	Turn        uint16
	Initialized bool
}

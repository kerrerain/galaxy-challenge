package game

import (
	"math"
)

// Computes the median distance between a planet and other planets
//
func (m Map) ComputeMedianDistance(planetID uint16, otherPlanetsID []uint16) Distance {
	if len(otherPlanetsID) == 0 {
		return Distance{
			Raw:   0.0,
			Turns: 0,
		}
	}

	medianRaw := 0.0
	medianTurns := uint16(0)
	length := uint16(len(otherPlanetsID))

	for _, otherPlanetID := range otherPlanetsID {
		medianRaw += m.DistanceMap[planetID][otherPlanetID].Raw
		medianTurns += m.DistanceMap[planetID][otherPlanetID].Turns
	}

	return Distance{
		Raw:   medianRaw / float64(length),
		Turns: uint16(math.Ceil(float64(medianTurns / uint16(length)))),
	}
}

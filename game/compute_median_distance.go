package game

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"math"
)

// Computes the median distance between a planet and enemy planets
//
func (m Map) ComputeMedianDistance(planetID uint16) Distance {
	otherPlanets := dto.FilterStatusPlanets(m.Planets, func(planet dto.StatusPlanet) bool {
		return planet.OwnerID != common.PLAYER_OWNER_ID && planet.OwnerID != common.NEUTRAL_OWNER_ID
	})

	if len(otherPlanets) == 0 {
		return Distance{
			Raw:   0.0,
			Turns: 0,
		}
	}

	medianRaw := 0.0
	medianTurns := uint16(0)
	length := uint16(len(otherPlanets))

	for _, planet := range otherPlanets {
		medianRaw += m.DistanceMap[planetID][planet.ID].Raw
		medianTurns += m.DistanceMap[planetID][planet.ID].Turns
	}

	return Distance{
		Raw:   medianRaw / float64(length),
		Turns: uint16(math.Ceil(float64(medianTurns / uint16(length)))),
	}
}

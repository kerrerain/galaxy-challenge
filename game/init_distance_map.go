package game

import (
	"errors"
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"math"
)

func (m *Map) InitDistanceMap() error {
	if m.Planets == nil {
		return errors.New("The planets have not been initialized.")
	}

	m.DistanceMap = make(map[int16]map[int16]Distance)

	for _, planet := range m.Planets {
		m.DistanceMap[planet.ID] = initDistanceMapForPlanet(planet, m.Planets)
	}

	return nil
}

func initDistanceMapForPlanet(planet dto.StatusPlanet, otherPlanets []dto.StatusPlanet) map[int16]Distance {
	planetDistanceMap := make(map[int16]Distance)

	for _, otherPlanet := range otherPlanets {
		rawDistance := computeDistance(planet, otherPlanet)
		distanceInTurns := computeTurnsLeft(rawDistance)

		planetDistanceMap[otherPlanet.ID] = Distance{
			Raw:   rawDistance,
			Turns: distanceInTurns,
		}
	}

	return planetDistanceMap
}

func computeDistance(p1 dto.StatusPlanet, p2 dto.StatusPlanet) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}

func computeTurnsLeft(rawDistance float64) int16 {
	return int16(math.Floor(rawDistance / common.DISTANCE_PER_TURN))
}

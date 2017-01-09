package paimon

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"math"
	"sort"
)

func evaluate(gameMap *game.Map) []*PlanetEvaluation {
	result := make([]*PlanetEvaluation, 0)

	ownPlanets := dto.FilterStatusPlanets(gameMap.Planets, func(planet dto.StatusPlanet) bool {
		return planet.OwnerID == common.PLAYER_OWNER_ID
	})

	for _, otherPlanet := range gameMap.Planets {
		result = append(result, &PlanetEvaluation{
			Planet: otherPlanet,
			Score:  evaluateDistance(gameMap, ownPlanets, otherPlanet),
		})
	}

	sort.Sort(ByGreatestScore(result))

	return result
}

func evaluateBasicStatus(planet dto.StatusPlanet) int {
	return 10000*ownerEvaluation(planet.OwnerID) +
		1000*int(planet.Growth) +
		1*int(planet.MaxUnits)
}

func ownerEvaluation(ownerID uint16) int {
	if ownerID == common.NEUTRAL_OWNER_ID {
		return 3
	} else if ownerID == common.PLAYER_OWNER_ID {
		return 1
	} else {
		return 2
	}
}

func evaluateDistance(gameMap *game.Map, ownPlanets []dto.StatusPlanet, otherPlanet dto.StatusPlanet) int {
	var minDistance float64

	for _, planet := range ownPlanets {
		distance := gameMap.DistanceMap[planet.ID][otherPlanet.ID]

		if planet.ID != otherPlanet.ID && minDistance > distance {
			minDistance = distance
		}
	}

	if minDistance == 0 {
		return 0
	} else {
		return int(math.Ceil(100 / minDistance * 1000000))
	}
}

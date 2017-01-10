package paimon

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"math"
	"sort"
)

type Evaluation struct {
	Map *game.Map
}

func CreateEvaluation(gameMap *game.Map) *Evaluation {
	return &Evaluation{
		Map: gameMap,
	}
}

func (e Evaluation) Run() []*PlanetEvaluation {
	result := make([]*PlanetEvaluation, 0)

	for _, otherPlanet := range e.Map.Planets {
		if otherPlanet.OwnerID != common.PLAYER_OWNER_ID {
			result = append(result, &PlanetEvaluation{
				Planet: otherPlanet,
				Score:  evaluateBasicStatus(otherPlanet) + e.evaluateDistance(otherPlanet),
			})
		}
	}

	sort.Sort(ByGreatestScore(result))

	return result
}

func (e Evaluation) evaluateDistance(otherPlanet dto.StatusPlanet) int {
	var minDistance float64

	ownPlanets := dto.FilterStatusPlanets(e.Map.Planets, func(planet dto.StatusPlanet) bool {
		return planet.OwnerID == common.PLAYER_OWNER_ID
	})

	for _, planet := range ownPlanets {
		distance := e.Map.DistanceMap[planet.ID][otherPlanet.ID]

		if planet.ID != otherPlanet.ID && (minDistance == 0 || distance < minDistance) {
			minDistance = distance
		}
	}

	if minDistance == 0 {
		return 0
	} else {
		return int(math.Ceil(100 / minDistance * 1000000))
	}
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

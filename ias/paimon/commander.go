package paimon

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"sort"
)

type Commander struct {
	Map                       *game.Map
	UnitsAlreadySentForPlanet map[uint16]int16
}

func CreateCommander(gameMap *game.Map) *Commander {
	return &Commander{
		Map: gameMap,
		UnitsAlreadySentForPlanet: make(map[uint16]int16),
	}
}

func (c *Commander) PlanInvasion(evaluations []*PlanetEvaluation) dto.Move {
	move := dto.CreateMove()

	for _, evaluation := range evaluations {
		move.Fleets = append(move.Fleets, c.planetInvasion(evaluation.Planet)...)
	}

	return *move
}

func (c *Commander) planetInvasion(targetPlanet dto.StatusPlanet) []dto.MoveFleet {
	fleets := make([]dto.MoveFleet, 0)
	unitsNeeded := targetPlanet.Units + 1
	distancesFromSourcePlanets := c.distancesFromSourcePlanets(targetPlanet)

	// Take all available units from nearby planets
	for _, distance := range distancesFromSourcePlanets {
		possibleUnits := c.possibleUnits(unitsNeeded, distance.Planet)

		if unitsNeeded > 0 && possibleUnits > 3 {
			fleets = append(fleets, dto.MoveFleet{
				SourceID: distance.Planet.ID,
				TargetID: targetPlanet.ID,
				Units:    possibleUnits,
			})

			unitsNeeded = unitsNeeded - possibleUnits
			c.UnitsAlreadySentForPlanet[distance.Planet.ID] = c.UnitsAlreadySentForPlanet[distance.Planet.ID] + possibleUnits
		}
	}

	return fleets
}

func (c *Commander) distancesFromSourcePlanets(targetPlanet dto.StatusPlanet) []Distance {
	sourcePlanets := dto.FilterStatusPlanets(c.Map.Planets, func(sourcePlanet dto.StatusPlanet) bool {
		return sourcePlanet.OwnerID == common.PLAYER_OWNER_ID
	})

	distances := make([]Distance, len(sourcePlanets))

	for index, sourcePlanet := range sourcePlanets {
		distances[index] = Distance{
			Distance: c.Map.DistanceMap[sourcePlanet.ID][targetPlanet.ID],
			Planet:   sourcePlanet,
		}
	}

	sort.Sort(ByLowestDistance(distances))

	return distances
}

func (c *Commander) possibleUnits(unitsNeeded int16, planet dto.StatusPlanet) int16 {
	return common.Max(0, common.Min(planet.Units-c.UnitsAlreadySentForPlanet[planet.ID]-5, unitsNeeded))
}

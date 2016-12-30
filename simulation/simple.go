package simulation

import (
	"github.com/magleff/galaxy-challenge/globals"
	"github.com/magleff/galaxy-challenge/models/game"
	"github.com/magleff/galaxy-challenge/models/move"
	"sort"
)

func SimulateSimple(g *game.Game) move.Move {
	// Full attack (full attack from a combination of the closest planets)
	// Reasonable attack (only 50% of the forces spent)
	// Full defense (focusing on defending planets under attack)
	fullAttackSimulations := make([]SimpleSimulation, 0)

	for _, planet := range g.Planets {
		fullAttackSimulations = append(fullAttackSimulations, simulateFullAttackOnPlanet(planet, g))
	}

	sort.Sort(ByGreatestFinalScore(fullAttackSimulations))

	// TODO sort by score
	if len(fullAttackSimulations) > 0 {
		return fullAttackSimulations[0].Move
	} else {
		return move.CreateMove()
	}
}

func simulateFullAttackOnPlanet(planet *game.Planet, g *game.Game) SimpleSimulation {
	simulation := SimpleSimulation{
		PlanetsFuture: initPlanetsFuture(g.Planets),
		Move:          CreateFutureMove(planet),
		FinalScore:    0,
	}

	simulation.Run(g, globals.HORIZON)

	return simulation
}

func initPlanetsFuture(planets []*game.Planet) []*PlanetFuture {
	planetsFuture := make([]*PlanetFuture, 0)

	for _, planet := range planets {
		planetsFuture = append(planetsFuture, &PlanetFuture{
			ID:       planet.ID,
			Units:    planet.Units,
			MaxUnits: planet.MaxUnits,
			Growth:   planet.Growth,
			OwnerID:  planet.OwnerID,
		})
	}

	return planetsFuture
}

func CreateFutureMove(targetPlanet *game.Planet) move.Move {
	futureMove := move.CreateMove()
	unitsNeeded := targetPlanet.Units + 1

	// Take all available units from nearby planets
	for _, distance := range targetPlanet.DistancesSorted {
		if distance.Planet.OwnerID == globals.PLAYER_OWNER_ID &&
			unitsNeeded > 0 &&
			distance.Planet.Units-1 > 3 {

			futureMove.Fleets = append(futureMove.Fleets, move.Fleet{
				SourceID: distance.Planet.ID,
				TargetID: targetPlanet.ID,
				Units:    Max(3, Min(distance.Planet.Units-1, unitsNeeded)),
			})

			unitsNeeded = unitsNeeded - Min(distance.Planet.Units-1, unitsNeeded)
		}
	}

	return futureMove
}

func Min(x, y int16) int16 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int16) int16 {
	if x > y {
		return x
	}
	return y
}

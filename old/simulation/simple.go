package simulation

import (
	"github.com/magleff/galaxy-challenge/globals"
	"github.com/magleff/galaxy-challenge/models/game"
	"github.com/magleff/galaxy-challenge/models/move"
	"log"
	"sort"
)

type SimulatedMove struct {
	Move  *move.Move
	Depth int
}

func SimulateSimple(g *game.Game) *move.Move {
	simulatedMove := &SimulatedMove{
		Move:  move.CreateMove(),
		Depth: 0,
	}

	fullAttackSimulation(simulatedMove, g)

	return simulatedMove.Move
}

// Full attack (full attack from a combination of the closest planets)
// Reasonable attack (only 50% of the forces spent)
// Full defense (focusing on defending planets under attack)
func fullAttackSimulation(simulatedMove *SimulatedMove, g *game.Game) {
	simulatedMove.Depth = simulatedMove.Depth + 1
	fullAttackSimulations := make([]SimpleSimulation, 0)

	for _, planet := range g.Planets {
		fullAttackSimulations = append(fullAttackSimulations, simulateFullAttackOnPlanet(planet, g))
	}

	sort.Sort(ByGreatestFinalScore(fullAttackSimulations))

	if len(fullAttackSimulations) > 0 {
		log.Println("Best move", fullAttackSimulations[0].FinalScore, fullAttackSimulations[0].Move)
	}

	if len(fullAttackSimulations) > 0 &&
		len(fullAttackSimulations[0].Move.Fleets) > 0 &&
		simulatedMove.Depth < globals.MAX_DEPTH {

		removeUnitsFromSourcePlanets(g, fullAttackSimulations[0].Move)
		//
		addMoveToGlobalMove(simulatedMove.Move, fullAttackSimulations[0].Move)
		// Recursive until there is no more fleet sent
		fullAttackSimulation(simulatedMove, g)
	} else {
		log.Println("Depth", simulatedMove.Depth)
	}
}

func removeUnitsFromSourcePlanets(g *game.Game, futureMove move.Move) {
	for _, fleet := range futureMove.Fleets {
		g.PlanetsByID[fleet.SourceID].Units = g.PlanetsByID[fleet.SourceID].Units - fleet.Units
	}
}

func addMoveToGlobalMove(globalMove *move.Move, fullAttackMove move.Move) {
	globalMove.Fleets = append(globalMove.Fleets, fullAttackMove.Fleets...)
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
	futureMove := *move.CreateMove()
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

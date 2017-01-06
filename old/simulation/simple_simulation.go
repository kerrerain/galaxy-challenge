package simulation

import (
	"github.com/magleff/galaxy-challenge/globals"
	"github.com/magleff/galaxy-challenge/models/game"
	"github.com/magleff/galaxy-challenge/models/move"
	"github.com/magleff/galaxy-challenge/models/status"
	"math"
)

type SimpleSimulation struct {
	PlanetsFuture []*PlanetFuture
	Move          move.Move
	FinalScore    int16
}

type ByGreatestFinalScore []SimpleSimulation

func (a ByGreatestFinalScore) Len() int           { return len(a) }
func (a ByGreatestFinalScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByGreatestFinalScore) Less(i, j int) bool { return a[i].FinalScore > a[j].FinalScore }

func (s *SimpleSimulation) Run(g *game.Game, horizon int) {
	moveFleets := createFleetsFromMove(g, s.Move)

	for _, planet := range s.PlanetsFuture {
		planet.FleetsArrivalsForPlanet(moveFleets, g.Fleets)
	}

	for i := 0; i < horizon; i++ {
		for _, planet := range s.PlanetsFuture {
			planet.Grow()
			planet.SimulateFleetsArrivals(i + 1)
		}
	}

	s.FinalScore = computeFinalScore(s.PlanetsFuture)
}

func computeFinalScore(planets []*PlanetFuture) int16 {
	finalScore := int16(0)

	for _, planet := range planets {
		if planet.OwnerID == globals.PLAYER_OWNER_ID {
			finalScore = finalScore + planet.Units
		}
	}

	return finalScore
}

func createFleetsFromMove(g *game.Game, futureMove move.Move) []status.Fleet {
	fleets := make([]status.Fleet, 0)

	for _, fleet := range futureMove.Fleets {
		fleets = append(fleets, status.Fleet{
			SourceID: fleet.SourceID,
			TargetID: fleet.TargetID,
			Units:    fleet.Units,
			Left:     computeTurnsLeft(g.DistanceFromTo(fleet.SourceID, fleet.TargetID)),
		})
	}

	return fleets
}

func computeTurnsLeft(distance float64) uint16 {
	return uint16(math.Ceil(distance / globals.DISTANCE_PER_TURN))
}

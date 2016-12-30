package simulation

import (
	"github.com/magleff/galaxy-challenge/models/game"
	"github.com/magleff/galaxy-challenge/models/move"
	"github.com/magleff/galaxy-challenge/models/status"
	"testing"
)

func TestRun(t *testing.T) {
	// Arrange
	var testCases = []struct {
		InputGame       *game.Game
		InputSimulation SimpleSimulation
		InputHorizon    int
		Expected        SimpleSimulation
	}{
		{
			&game.Game{
				PlanetsByID: map[uint16]*game.Planet{
					1: &game.Planet{
						Distances: map[uint16]*game.Distance{
							2: &game.Distance{Distance: 40},
						},
					},
				},
				Fleets: []status.Fleet{
					{OwnerID: 2, Units: 20, SourceID: 34, TargetID: 2, Left: 3},
				},
			},
			SimpleSimulation{
				PlanetsFuture: []*PlanetFuture{
					{ID: 1, OwnerID: 1, Units: 50, MaxUnits: 200, Growth: 5}, // Own planet
					{ID: 2, OwnerID: 0, Units: 30, MaxUnits: 100, Growth: 3}, // Neutral planet
				},
				Move: move.Move{
					Fleets: []move.Fleet{
						{SourceID: 1, TargetID: 2, Units: 40},
					},
				},
			},
			4,
			SimpleSimulation{
				PlanetsFuture: []*PlanetFuture{
					{ID: 1, OwnerID: 1, Units: 30, MaxUnits: 200, Growth: 5},
					{ID: 2, OwnerID: 2, Units: 10, MaxUnits: 100, Growth: 3,
						FleetsArrivals: map[int][]status.Fleet{
							3: []status.Fleet{
								{OwnerID: 2, Units: 20, SourceID: 34, TargetID: 2, Left: 3},
							},
							2: []status.Fleet{
								{OwnerID: 1, Units: 40, SourceID: 1, TargetID: 2, Left: 2},
							},
						}}, // The enemy took back the neutral planet
				},
				Move: move.Move{
					Fleets: []move.Fleet{
						{SourceID: 1, TargetID: 2, Units: 40},
					},
				},
				FinalScore: 30, // Our player lost units
			},
		},
	}

	for index, testCase := range testCases {
		// Arrange
		simulation := testCase.InputSimulation

		// Act
		simulation.Run(testCase.InputGame, testCase.InputHorizon)

		// Assert
		if testCase.Expected.FinalScore != simulation.FinalScore {
			t.Errorf("TestRun(%d): expected %s, actual %s", index, testCase.Expected.FinalScore, simulation.FinalScore)
		}
	}
}

func TestComputeTurnsLeft(t *testing.T) {
	// Arrange
	var testCases = []struct {
		Input    float64
		Expected uint16
	}{
		{
			Input:    40,
			Expected: 2,
		},
		{
			Input:    45,
			Expected: 3,
		},
	}

	for index, testCase := range testCases {
		// Act
		actual := computeTurnsLeft(testCase.Input)

		// Assert
		if actual != testCase.Expected {
			t.Errorf("TestRun(%d): expected %s, actual %s", index, testCase.Expected, actual)
		}
	}
}

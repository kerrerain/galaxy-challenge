package game

import (
	"reflect"
	"testing"
)

func TestFilterPlanets(t *testing.T) {
	// Arrange
	var testCases = []struct {
		Input     []Planet
		Predicate func(Planet) bool
		Expected  []Planet
	}{
		{
			[]Planet{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			func(planet Planet) bool {
				return planet.ID != 2
			},
			[]Planet{
				{
					ID: 1,
				},
			},
		},
	}

	for index, testCase := range testCases {
		// Act
		actual := FilterPlanets(testCase.Input, testCase.Predicate)

		// Assert
		if !reflect.DeepEqual(actual, testCase.Expected) {
			t.Errorf("TestFilterPlanets(%d): expected %s, actual %s", index, testCase.Expected, actual)
		}
	}
}

func TestInitializeDistances(t *testing.T) {
	var testCases = []struct {
		Input                       []*Planet
		ExpectedDistanceToPlanet3   float64
		ExpectedMinDistancePlanetID uint16
	}{
		{
			[]*Planet{
				{
					ID: 1,
					X:  0,
					Y:  0,
				},
				{
					ID: 2,
					X:  1,
					Y:  0,
				},
				{
					ID: 3,
					X:  3,
					Y:  0,
				},
			},
			3,
			2,
		},
	}

	for index, testCase := range testCases {
		// Arrange
		planet := &Planet{
			ID:        1,
			X:         0,
			Y:         0,
			Distances: make(map[uint16]*Distance),
		}

		// Act
		planet.InitializeDistances(testCase.Input)

		// Assert
		if testCase.ExpectedDistanceToPlanet3 != planet.Distances[3].Distance {
			t.Errorf("TestInitializeDistances(%d): expected %s, actual %s", index,
				testCase.ExpectedDistanceToPlanet3, planet.Distances[3].Distance)
		}

		if testCase.ExpectedMinDistancePlanetID != planet.DistancesSorted[0].Planet.ID {
			t.Errorf("TestInitializeDistances(%d): expected %s, actual %s", index,
				testCase.ExpectedMinDistancePlanetID, planet.DistancesSorted[0].Planet.ID)
		}
	}
}

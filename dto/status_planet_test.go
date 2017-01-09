package dto

import (
	"reflect"
	"testing"
)

func TestFilterPlanets(t *testing.T) {
	// Arrange
	var testCases = []struct {
		Input     []StatusPlanet
		Predicate func(StatusPlanet) bool
		Expected  []StatusPlanet
	}{
		{
			[]StatusPlanet{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			func(planet StatusPlanet) bool {
				return planet.ID != 2
			},
			[]StatusPlanet{
				{
					ID: 1,
				},
			},
		},
	}

	for index, testCase := range testCases {
		// Act
		actual := FilterStatusPlanets(testCase.Input, testCase.Predicate)

		// Assert
		if !reflect.DeepEqual(actual, testCase.Expected) {
			t.Errorf("TestFilterPlanets(%d): expected %s, actual %s", index, testCase.Expected, actual)
		}
	}
}

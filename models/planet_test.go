package models

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

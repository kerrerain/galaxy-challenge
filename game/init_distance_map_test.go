package game

import (
	"github.com/magleff/galaxy-challenge/dto"
	"reflect"
	"testing"
)

func TestInitDistanceMapEmptyPlanets(t *testing.T) {
	// Arrange
	gameMap := &Map{}

	// Act
	err := gameMap.InitDistanceMap()

	// Assert
	if err == nil {
		t.Error("TestInitDistanceMapEmptyPlanets: should trigger an error if the planets map is empty.")
	}
}

func TestInitDistanceMap(t *testing.T) {
	// Arrange
	var testCases = []struct {
		Input             []dto.StatusPlanet
		InputSourcePlanet uint16
		InputTargetPlanet uint16
		Expected          Distance
	}{
		{
			[]dto.StatusPlanet{
				{ID: 1, X: 0, Y: 0},
				{ID: 2, X: 15, Y: 0},
				{ID: 3, X: 5, Y: 0},
			},
			1, 3, Distance{
				Raw:   5,
				Turns: 0,
			},
		},
	}

	for index, testCase := range testCases {
		// Arrange
		gameMap := &Map{
			Planets: testCase.Input,
		}

		// Act
		gameMap.InitDistanceMap()

		// Assert
		actual := gameMap.DistanceMap[testCase.InputSourcePlanet][testCase.InputTargetPlanet]

		if !reflect.DeepEqual(testCase.Expected, actual) {
			t.Errorf("TestInitDistanceMap(%d): expected %s, was %s", index,
				testCase.Expected, actual)
		}
	}
}

func TestComputeTurnsLeft(t *testing.T) {
	// Arrange
	gameMap := &Map{
		Planets: []dto.StatusPlanet{
			{ID: 1, X: 0, Y: 0},
			{ID: 2, X: 40, Y: 0},
			{ID: 3, X: 45, Y: 0},
		},
	}

	gameMap.InitDistanceMap()

	testCases := []struct {
		rawDistance float64
		Expected    uint16
	}{
		{40, 2},
		{45, 2},
		{5, 0},
	}

	for index, testCase := range testCases {
		// Act
		actual := computeTurnsLeft(testCase.rawDistance)

		// Assert
		if actual != testCase.Expected {
			t.Errorf("TestComputeTurnsLeft(%d): expected %d, was %d", index, testCase.Expected, actual)
		}
	}
}

package paimon

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"reflect"
	"testing"
)

func TestPlanInvasion(t *testing.T) {
	// Arrange
	planets := []dto.StatusPlanet{
		{ID: 1, X: 0, Y: 15, OwnerID: common.PLAYER_OWNER_ID, Units: 100, MaxUnits: 200},
		{ID: 2, X: 0, Y: 10, OwnerID: common.PLAYER_OWNER_ID, Units: 40, MaxUnits: 80},
		{ID: 3, X: 0, Y: 5, OwnerID: common.NEUTRAL_OWNER_ID, Units: 50, MaxUnits: 200},
	}

	testCases := []struct {
		InputPlanets   []*PlanetEvaluation
		ExpectedFleets []dto.MoveFleet
	}{
		// Basic test case
		{
			[]*PlanetEvaluation{
				{
					Score:  1000,
					Planet: planets[2],
				},
			},
			[]dto.MoveFleet{
				{
					SourceID: 2,
					TargetID: 3,
					Units:    35,
				},
				{
					SourceID: 1,
					TargetID: 3,
					Units:    16,
				},
			},
		},
	}

	for index, testCase := range testCases {
		// Arrange
		gameMap := &game.Map{
			Planets: planets,
		}
		gameMap.InitDistanceMap()
		commander := CreateCommander(gameMap)

		// Act
		move := commander.PlanInvasion(testCase.InputPlanets)

		// Assert
		if !reflect.DeepEqual(move.Fleets, testCase.ExpectedFleets) {
			t.Errorf("TestPlanInvasion (%d): the fleets sent are not correct.", index)
		}
	}
}

func TestPossibleUnits(t *testing.T) {
	// Arrange
	testCases := []struct {
		InputUnitsAlreadySent map[uint16]int16
		InputSourcePlanet     dto.StatusPlanet
		InputUnitsNeeded      int16
		Expected              int16
	}{
		// Case 1: There are no units already taken,
		// but the minimum of units is set to 5.
		{
			map[uint16]int16{
				1: 0,
			},
			dto.StatusPlanet{
				ID:    1,
				Units: 50,
			},
			50,
			45,
		},
		{
			map[uint16]int16{
				1: 0,
			},
			dto.StatusPlanet{
				ID:    1,
				Units: 50,
			},
			20,
			20,
		},
		{
			map[uint16]int16{
				1: 20,
			},
			dto.StatusPlanet{
				ID:    1,
				Units: 50,
			},
			50,
			25,
		},
		{
			map[uint16]int16{
				1: 50,
			},
			dto.StatusPlanet{
				ID:    1,
				Units: 50,
			},
			50,
			0,
		},
	}

	for index, testCase := range testCases {
		// Arrange
		commander := CreateCommander(&game.Map{})
		commander.UnitsAlreadySentForPlanet = testCase.InputUnitsAlreadySent

		// Act
		actual := commander.possibleUnits(testCase.InputUnitsNeeded, testCase.InputSourcePlanet)

		// Assert
		if testCase.Expected != actual {
			t.Errorf("TestPossibleUnits (%d): expected %d, was %d.", index, testCase.Expected, actual)
		}
	}
}

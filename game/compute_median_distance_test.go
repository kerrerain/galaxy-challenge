package game

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"reflect"
	"testing"
)

func TestComputeMedianDistance(t *testing.T) {
	// Arrange
	testCases := []struct {
		Input    []dto.StatusPlanet
		Expected Distance
	}{
		{
			[]dto.StatusPlanet{},
			Distance{
				Raw:   0,
				Turns: 0,
			},
		},
		{
			[]dto.StatusPlanet{
				{ID: 1, OwnerID: common.PLAYER_OWNER_ID, X: 0, Y: 0},
				{ID: 2, OwnerID: common.NEUTRAL_OWNER_ID, X: 45, Y: 0},
				{ID: 3, OwnerID: 2, X: 40, Y: 0},
				{ID: 4, OwnerID: 3, X: 80, Y: 0},
			},
			Distance{
				Raw:   60,
				Turns: 3,
			},
		},
	}

	for i, testCase := range testCases {
		// Arrange
		gameMap := &Map{
			Planets: testCase.Input,
		}
		gameMap.InitDistanceMap()

		// Act
		actual := gameMap.ComputeMedianDistance(1, []int16{3, 4})

		// Assert
		if !reflect.DeepEqual(testCase.Expected, actual) {
			t.Errorf("TestComputeMedianDistance(%d): expected %d, was %d", i, testCase.Expected, actual)
		}
	}
}

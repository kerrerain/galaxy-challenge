package command

import (
	"github.com/magleff/galaxy-challenge/dto"
	"reflect"
	"testing"
)

func TestBuildMove(t *testing.T) {
	// Arrange
	testCases := []struct {
		Orders       []Order
		ExpectedMove dto.Move
	}{
		{
			[]Order{
				{
					TargetID: 2,
					SourceID: 1,
					Units:    5,
				},
				{
					TargetID: 2,
					SourceID: 1,
					Units:    5,
				},
			},
			dto.Move{
				Fleets: []dto.MoveFleet{
					{SourceID: 1, TargetID: 2, Units: 10},
				},
			},
		},
	}

	for index, testCase := range testCases {
		// Arrange
		commander := &Commander{
			AvailableUnitsOnPlanet: map[uint16]int16{1: 50},
		}

		// Act
		for _, order := range testCase.Orders {
			commander.SendOrder(order)
		}

		actual := commander.BuildMove()

		// Assert
		if !reflect.DeepEqual(actual, testCase.ExpectedMove) {
			t.Errorf("TestBuildMove(%d): Expected (%v) but was (%v)", index,
				testCase.ExpectedMove, actual)
		}
	}
}

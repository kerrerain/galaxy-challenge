package command

import (
	"reflect"
	"testing"
)

func TestSendOrder(t *testing.T) {
	// Arrange
	testCases := []struct {
		Input    Order
		Expected map[int16]int16
	}{
		// Case 1: a single source planet
		{
			Order{
				TargetID: 2,
				SourceID: 1,
				Units:    5,
			},
			map[int16]int16{1: 5},
		},
		// Case 2: not enough units to send (a fleet should be >= 3 units)
		{
			Order{
				TargetID: 2,
				SourceID: 1,
				Units:    2,
			},
			map[int16]int16{1: 10},
		},
		// Case 3: not enough units available.
		// Send possible units but keep at least 1 unit
		{
			Order{
				TargetID: 2,
				SourceID: 1,
				Units:    15,
			},
			map[int16]int16{1: 1},
		},
	}

	for index, testCase := range testCases {
		// Arrange
		commander := &Commander{
			AvailableUnitsOnPlanet: map[int16]int16{1: 10},
		}

		// Act
		commander.SendOrder(testCase.Input)

		// Assert
		if !reflect.DeepEqual(commander.AvailableUnitsOnPlanet, testCase.Expected) {
			t.Errorf("TestSendOrder(%d): Expected (%v) but was (%v)", index,
				testCase.Expected, commander.AvailableUnitsOnPlanet)
		}
	}
}

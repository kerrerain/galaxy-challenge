package command

import (
	"testing"
)

func TestUnitsToSend(t *testing.T) {
	// Arrange
	testCases := []struct {
		AvailableUnits    int16
		AttackUnits       int16
		ExpectedUnitsSent int16
	}{
		// Case 1: Enough units to send.
		{10, 5, 5},

		// Case 2: Can send at most (available - 1) units.
		// There must be at least 1 unit left.
		{10, 15, 9},

		// Case 3: Cannot send less than 3 units.
		{10, 2, 0},
	}

	for index, testCase := range testCases {
		// Act
		actual := unitsToSend(testCase.AvailableUnits, testCase.AttackUnits)

		// Assert
		if actual != testCase.ExpectedUnitsSent {
			t.Errorf("TestUnitsToSend(%d): Expected (%d) but was (%d)", index, testCase.ExpectedUnitsSent, actual)
		}
	}
}

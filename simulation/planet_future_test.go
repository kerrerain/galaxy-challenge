package simulation

import (
	"github.com/magleff/galaxy-challenge/globals"
	"github.com/magleff/galaxy-challenge/models/status"
	"testing"
)

func TestGrow(t *testing.T) {
	// Arrange
	var testCases = []struct {
		Input    *PlanetFuture
		Expected int16
	}{
		{
			Input: &PlanetFuture{
				OwnerID:  globals.PLAYER_OWNER_ID,
				Growth:   5,
				Units:    10,
				MaxUnits: 20,
			},
			Expected: 15,
		},
		{
			Input: &PlanetFuture{
				OwnerID:  globals.PLAYER_OWNER_ID,
				Growth:   5,
				Units:    18,
				MaxUnits: 20,
			},
			Expected: 20,
		},
		{
			Input: &PlanetFuture{
				OwnerID:  globals.NEUTRAL_OWNER_ID,
				Growth:   5,
				Units:    10,
				MaxUnits: 20,
			},
			Expected: 10,
		},
	}

	for index, testCase := range testCases {
		// Arrange
		planetFuture := testCase.Input

		// Act
		planetFuture.Grow()

		// Assert
		if planetFuture.Units != testCase.Expected {
			t.Errorf("TestRun(%d): expected %s, actual %s", index, testCase.Expected, planetFuture.Units)
		}
	}
}

func TestSimulateFleetsArrivals(t *testing.T) {
	// Arrange
	var testCases = []struct {
		Input    *PlanetFuture
		Expected int16
	}{
		{
			Input: &PlanetFuture{
				Units: 10,
				FleetsArrivals: map[int][]status.Fleet{
					1: []status.Fleet{
						{Units: 20},
					},
				},
			},
			Expected: -10,
		},
	}

	for index, testCase := range testCases {
		// Arrange
		planetFuture := testCase.Input

		// Act
		planetFuture.SimulateFleetsArrivals(1)

		// Assert
		if planetFuture.Units != testCase.Expected {
			t.Errorf("TestRun(%d): expected %s, actual %s", index, testCase.Expected, planetFuture.Units)
		}
	}
}

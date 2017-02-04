package evaluation

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/engine"
	"testing"
)

func TestComputeGrowth(t *testing.T) {
	// Arrange
	testCases := []struct {
		Input    *engine.PlanetTimeline
		Expected int16
	}{
		// Case 1: the player owns the planet
		{
			&engine.PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 100, MaxUnits: 200},
					{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 125, MaxUnits: 200},
					{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 150, MaxUnits: 200},
				},
			},
			25,
		},
	}

	for index, testCase := range testCases {
		// Arrange
		planetTimeline := testCase.Input

		// Act
		actual := ComputeGrowth(planetTimeline)

		// Assert
		if actual != testCase.Expected {
			t.Errorf("TestComputeGrowth(%d): expected (%d), actual was (%d)", index, testCase.Expected, actual)
		}
	}
}

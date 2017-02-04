package evaluation

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/engine"
	"testing"
)

func TestComputeLoss(t *testing.T) {
	// Arrange
	testCases := []struct {
		Input    *engine.PlanetTimeline
		Expected int
	}{
		// Case 1: the player has lost the planet
		{
			&engine.PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 100, MaxUnits: 200},
					{ID: 1, OwnerID: 2, Units: 50, MaxUnits: 200},
				},
			},
			-1,
		},
		// Case 2: the player keeps the planet
		{
			&engine.PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 100, MaxUnits: 200},
					{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 150, MaxUnits: 200},
				},
			},
			0,
		},
		// Case 3: the player earns the planet
		{
			&engine.PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{ID: 1, OwnerID: common.NEUTRAL_OWNER_ID, Units: 100, MaxUnits: 200},
					{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 150, MaxUnits: 200},
				},
			},
			1,
		},
		// Case 4: nothing to compute
		{
			&engine.PlanetTimeline{
				Turns: []dto.StatusPlanet{},
			},
			0,
		},
	}

	for index, testCase := range testCases {
		// Arrange
		planetTimeline := testCase.Input

		// Act
		actual := ComputeLoss(planetTimeline)

		// Assert
		if actual != testCase.Expected {
			t.Errorf("TestComputeOwnership(%d): expected (%d), actual was (%d)", index, testCase.Expected, actual)
		}
	}
}

package core

import (
	"github.com/magleff/galaxy-challenge/models"
	"reflect"
	"testing"
)

func TestTargetPlanets(t *testing.T) {
	// Arrange
	var testCases = []struct {
		Input    models.Universe
		Expected []models.Planet
	}{
		// Do nothing
		{
			models.Universe{
				Planets: []models.Planet{},
				Fleets:  []models.Fleet{},
			},
			[]models.Planet{},
		},
		// Only target planets that are not owned by the player
		{
			models.Universe{
				Planets: []models.Planet{
					{OwnerID: PLAYER_OWNER_ID}, {OwnerID: NEUTRAL_OWNER_ID}, {OwnerID: 2},
				},
				Fleets: []models.Fleet{},
			},
			[]models.Planet{
				{OwnerID: NEUTRAL_OWNER_ID}, {OwnerID: 2},
			},
		},
		// Only target planets that are habitable (H, K, L, M)
		// Other planets are inhabitable (D, J, N)
		{
			models.Universe{
				Planets: []models.Planet{
					{Category: "M"}, {Category: "H"}, {Category: "K"}, {Category: "L"},
					{Category: "D"}, {Category: "J"}, {Category: "N"},
				},
				Fleets: []models.Fleet{},
			},
			[]models.Planet{
				{Category: "M"}, {Category: "H"}, {Category: "K"}, {Category: "L"},
			},
		},
		// Only target planets that the player has not targeted yet with a fleet
		{
			models.Universe{
				Planets: []models.Planet{
					{ID: 1}, {ID: 2},
				},
				Fleets: []models.Fleet{
					{TargetID: 1, OwnerID: PLAYER_OWNER_ID}, {TargetID: 2, OwnerID: 3},
				},
			},
			[]models.Planet{
				{ID: 2},
			},
		},
	}

	for index, testCase := range testCases {
		// Act
		actual := targetPlanets(testCase.Input)

		// Assert
		if !reflect.DeepEqual(actual, testCase.Expected) {
			t.Errorf("TestTargetPlanets(%d): expected %s, actual %s", index, testCase.Expected, actual)
		}
	}
}

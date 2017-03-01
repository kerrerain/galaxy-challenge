package simulation

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"reflect"
	"testing"
)

func TestComputeNeeds(t *testing.T) {
	// Arrange
	testCases := []struct {
		GameMap  *game.Map
		Expected Needs
	}{
		{
			&game.Map{
				Planets: []dto.StatusPlanet{},
				Fleets:  []dto.StatusFleet{},
			},
			Needs{
				ID:             0,
				AvailableUnits: 0,
				SupportDemands: []SupportDemand{},
			},
		},
		{
			&game.Map{
				Planets: []dto.StatusPlanet{
					{
						ID:       1,
						Units:    5,
						OwnerID:  common.PLAYER_OWNER_ID,
						Growth:   3,
						MaxUnits: 50,
					},
				},
				Fleets: []dto.StatusFleet{
					{OwnerID: 2, Units: 5, Left: 4, TargetID: 1},
					{OwnerID: 2, Units: 5, Left: 5, TargetID: 1},
				},
			},
			Needs{
				ID:             1,
				AvailableUnits: 4,
				SupportDemands: []SupportDemand{},
			},
		},
		{
			&game.Map{
				Planets: []dto.StatusPlanet{
					{
						ID:       1,
						Units:    1,
						OwnerID:  common.PLAYER_OWNER_ID,
						Growth:   3,
						MaxUnits: 50,
					},
				},
				Fleets: []dto.StatusFleet{
					{OwnerID: 2, Units: 10, Left: 2, TargetID: 1},
					{OwnerID: 2, Units: 10, Left: 3, TargetID: 1},
				},
			},
			Needs{
				ID:             1,
				AvailableUnits: 0,
				SupportDemands: []SupportDemand{
					{NeededUnits: 4, Left: 2, PenaltyAfter: 3, UselessBefore: false},
					{NeededUnits: 7, Left: 3, PenaltyAfter: 3, UselessBefore: false},
				},
			},
		},
		{
			&game.Map{
				Planets: []dto.StatusPlanet{
					{
						ID:       1,
						Units:    48,
						OwnerID:  common.PLAYER_OWNER_ID,
						Growth:   3,
						MaxUnits: 50,
					},
				},
				Fleets: []dto.StatusFleet{
					{OwnerID: 2, Units: 60, Left: 2, TargetID: 1},
				},
			},
			Needs{
				ID:             1,
				AvailableUnits: 3,
				SupportDemands: []SupportDemand{
					{NeededUnits: 11, Left: 2, PenaltyAfter: 3, UselessBefore: true},
				},
			},
		},
		{
			&game.Map{
				Planets: []dto.StatusPlanet{
					{
						ID:       1,
						Units:    48,
						OwnerID:  common.PLAYER_OWNER_ID,
						Growth:   3,
						MaxUnits: 50,
					},
				},
				Fleets: []dto.StatusFleet{
					{OwnerID: 2, Units: 10, Left: 2, TargetID: 1},
				},
			},
			Needs{
				ID:             1,
				AvailableUnits: 43,
				SupportDemands: []SupportDemand{},
			},
		},
	}

	for index, testCase := range testCases {
		gameMap := testCase.GameMap
		gameMap.InitDistanceMap()

		// Act
		actual := ComputeNeeds(gameMap, 1)

		// Assert
		if !reflect.DeepEqual(testCase.Expected, actual) {
			t.Errorf("TestComputeNeeds(%d): expected %v, was %v", index, testCase.Expected, actual)
		}
	}
}

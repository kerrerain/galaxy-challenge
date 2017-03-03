package simulation

import (
	"github.com/magleff/galaxy-challenge/command"
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"log"
	"testing"
)

func TestComputeKept(t *testing.T) {
	// Arrange
	testCases := []struct {
		GameMap  *game.Map
		SourceID int16
		Order    command.Order
		Expected bool
	}{
		{
			&game.Map{
				Planets: []dto.StatusPlanet{
					{ID: 1, OwnerID: 1, Units: 10, Growth: 2, MaxUnits: 200},
				},
				Fleets: []dto.StatusFleet{
					{OwnerID: 2, Units: 3, TargetID: 1, Left: 1},
					{OwnerID: 2, Units: 5, TargetID: 1, Left: 1},
				},
			},
			1,
			command.Order{
				SourceID: 1,
				TargetID: 0,
				Units:    1,
			},
			true,
		},
		{
			&game.Map{
				Planets: []dto.StatusPlanet{
					{ID: 1, OwnerID: 1, Units: 10, Growth: 2, MaxUnits: 200},
				},
				Fleets: []dto.StatusFleet{
					{OwnerID: 2, Units: 3, TargetID: 1, Left: 1},
					{OwnerID: 2, Units: 5, TargetID: 1, Left: 1},
				},
			},
			1,
			command.Order{
				SourceID: 1,
				TargetID: 0,
				Units:    15,
			},
			false,
		},
	}

	for index, testCase := range testCases {
		if common.DEBUG_MODE {
			log.Printf("TestComputeKept(%d)", index)
		}

		gameMap := testCase.GameMap
		gameMap.InitDistanceMap()

		// Act
		actual := ComputeKept(gameMap, testCase.SourceID, testCase.Order, 1)

		// Assert
		if actual != testCase.Expected {
			t.Errorf("TestComputeKept(%d): expected %d, was %d", index, testCase.Expected, actual)
		}
	}
}

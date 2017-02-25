package engine

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"log"
	"reflect"
	"testing"
)

func TestCreateTimeline(t *testing.T) {
	// Arrange
	gameMap := &game.Map{
		Planets: []dto.StatusPlanet{
			{ID: 1},
		},
		Fleets: []dto.StatusFleet{
			{},
		},
	}

	// Act
	timeline := CreateTimeline(gameMap)

	// Assert
	if len(timeline.PlanetTimelines) != 1 {
		t.Error("TestCreateTimeline: should initialize the planets.")
	}
}

func TestTimelineNextTurn(t *testing.T) {
	// Arrange
	testCases := []struct {
		GameMap        *game.Map
		ExpectedPlanet dto.StatusPlanet
	}{
		{
			&game.Map{
				Planets: []dto.StatusPlanet{
					{ID: 1, OwnerID: common.NEUTRAL_OWNER_ID, Units: 50, MaxUnits: 200, Growth: 5, Category: "M"},
				},
				Fleets: []dto.StatusFleet{
					{OwnerID: common.PLAYER_OWNER_ID, TargetID: 1, Units: 60, Left: 0},
				},
			},
			dto.StatusPlanet{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 10, MaxUnits: 200, Growth: 5, Category: "M"},
		},
		{
			&game.Map{
				Planets: []dto.StatusPlanet{
					{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 50, MaxUnits: 200, Growth: 5, Category: "M"},
				},
				Fleets: []dto.StatusFleet{
					{OwnerID: 2, TargetID: 1, Units: 60, Left: 0},
				},
			},
			dto.StatusPlanet{ID: 1, OwnerID: 2, Units: 5, MaxUnits: 200, Growth: 5, Category: "M"},
		},
	}

	for index, testCase := range testCases {
		if common.DEBUG_MODE {
			log.Printf("Case %d", index)
		}

		timeline := CreateTimeline(testCase.GameMap)

		// Act
		timeline.NextTurn()

		// Assert
		if timeline.Turn != 1 {
			t.Error("TestTimelineNextTurn: should increase the turns.")
		}

		if !reflect.DeepEqual(testCase.ExpectedPlanet, timeline.PlanetTimelinesMap[1].CurrentTurn()) {
			t.Errorf("TestTimelineNextTurn(%d): expected %v, was %v", index, testCase.ExpectedPlanet,
				timeline.PlanetTimelinesMap[2].CurrentTurn())
		}
	}
}

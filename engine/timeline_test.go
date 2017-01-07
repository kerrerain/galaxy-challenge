package engine

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
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
	gameMap := &game.Map{
		Planets: []dto.StatusPlanet{
			{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 50, MaxUnits: 200, Growth: 5, Category: "M"},
			{ID: 2, OwnerID: common.NEUTRAL_OWNER_ID, Units: 50, MaxUnits: 200, Growth: 5, Category: "M"},
			{ID: 3, OwnerID: 2, Units: 50, MaxUnits: 200, Growth: 5, Category: "M"},
		},
		Fleets: []dto.StatusFleet{
			{},
		},
	}

	timeline := CreateTimeline(gameMap)

	// Act
	timeline.NextTurn()
	timeline.NextTurn()
	timeline.NextTurn()

	// Assert
	if timeline.Turn != 3 {
		t.Error("TestTimelineNextTurn: should increase the turns.")
	}
}

package game

/*import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"testing"
)

func TestEvaluateDefense(t *testing.T) {
	// Arrange
	gameMap := &Map{
		Planets: []*dto.StatusPlanet{
			{ID: 1, OwnerID: common.PLAYER_OWNER_ID, X: -100, Y: 0, Units: 200, Growth: 5},
			{ID: 2, OwnerID: common.PLAYER_OWNER_ID, X: -40, Y: 0, Units: 30, Growth: 3},
			{ID: 3, OwnerID: common.NEUTRAL_OWNER_ID, X: 0, Y: 0, Units: 40, Growth: 3},
			{ID: 4, OwnerID: 2, X: 40, Y: 0, Units: 30, Growth: 5},
			{ID: 5, OwnerID: 2, X: 100, Y: 0, Units: 200, Growth: 5},
		},
	}

	config := EvaluateDefenseConfig{Turns: 10, DistancePerTurn: 20}
	gameMap.InitDistanceMap()

	// Act
	gameMap.EvaluateDefense(config)

	// Assert
	actual := gameMap.Planets[2].DefensePotential

	if actual != -16 {
		t.Errorf("TestEvaluateDefense: expected -160, was %d", actual)
	}
}*/

package paimon

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
	"testing"
)

func TestEvaluationBasicStatus(t *testing.T) {
	// Arrange
	planets := []dto.StatusPlanet{
		// The planet is owned by the player. Without further analysis,
		// that's not the most interesting planet to send fleets.
		dto.StatusPlanet{
			OwnerID:  common.PLAYER_OWNER_ID,
			Growth:   5,
			MaxUnits: 200,
			Category: "M",
			Units:    30,
		},
		// The planet is owned by another player, so it will be harder to take.
		dto.StatusPlanet{
			OwnerID:  2,
			Growth:   5,
			MaxUnits: 200,
			Category: "M",
			Units:    30,
		},
		// The planet is very limited in growth. It is not interesting,
		// but it is still neutral.
		dto.StatusPlanet{
			OwnerID:  common.NEUTRAL_OWNER_ID,
			Growth:   1,
			MaxUnits: 30,
			Category: "J",
			Units:    30,
		},
		// The planet is neutral and is the best planet possible.
		// It should have the greatest score.
		dto.StatusPlanet{
			OwnerID:  common.NEUTRAL_OWNER_ID,
			Growth:   5,
			MaxUnits: 200,
			Category: "M",
			Units:    30,
		},
	}

	lastScore := 0

	for index, planet := range planets {
		// Act
		score := evaluateBasicStatus(planet)

		// Assert
		if score <= lastScore {
			t.Errorf("TestEvaluateBasicStatus: planet %d should not have such an high score: %d", index, score)
		}

		lastScore = score
	}
}

func TestEvaluateDistance(t *testing.T) {
	// Arrange
	planets := []dto.StatusPlanet{
		dto.StatusPlanet{
			ID:      1,
			OwnerID: common.PLAYER_OWNER_ID,
			X:       0,
			Y:       0,
		},
		dto.StatusPlanet{
			ID:      2,
			OwnerID: 2,
			X:       1,
			Y:       0,
		},
		dto.StatusPlanet{
			ID:      3,
			OwnerID: common.NEUTRAL_OWNER_ID,
			X:       3,
			Y:       0,
		},
		dto.StatusPlanet{
			ID:      4,
			OwnerID: common.NEUTRAL_OWNER_ID,
			X:       5,
			Y:       0,
		},
	}

	gameMap := &game.Map{
		Planets: planets,
	}
	gameMap.InitDistanceMap()

	evaluation := CreateEvaluation(gameMap)

	// Act
	score := evaluation.evaluateDistance(planets[2])

	// Assert
	if score != 300 {
		t.Errorf("TestEvaluateDistance: expected %d but was %d", 300, score)
	}
}

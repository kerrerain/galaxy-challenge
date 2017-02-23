package engine

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"testing"
)

func TestCreatePlanetTimeline(t *testing.T) {
	// Arrange
	planet := dto.StatusPlanet{}

	// Act
	planetTimeline := CreatePlanetTimeline(planet)

	// Assert
	if len(planetTimeline.Turns) != 1 {
		t.Error("TestCreatePlanetTimeline: should init the first turn (t=0) with the given status of the planet.")
	}
}

func TestNextTurn(t *testing.T) {
	// Arrange
	var testCases = []struct {
		Input                       *PlanetTimeline
		InputFleets                 []dto.StatusFleet
		ExpectedUnits               int16
		ExpectedOwnerID             int16
		ExpectedTotalUnitsSent      int16
		ExpectedTotalEnemyUnitsSent int16
	}{
		// CASE 1: normal growth
		{
			&PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{
						ID:       1,
						OwnerID:  common.PLAYER_OWNER_ID,
						Category: "M",
						Growth:   5,
						MaxUnits: 200,
						Units:    30,
					},
				},
			},
			[]dto.StatusFleet{},
			35,
			common.PLAYER_OWNER_ID,
			0,
			0,
		},
		// CASE 2: not growing because the planet is neutral
		{
			&PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{
						ID:       1,
						OwnerID:  common.NEUTRAL_OWNER_ID,
						Category: "M",
						Growth:   5,
						MaxUnits: 200,
						Units:    30,
					},
				},
			},
			[]dto.StatusFleet{},
			30,
			common.NEUTRAL_OWNER_ID,
			0,
			0,
		},
		// CASE 3: not growing because the growth rate is 0
		{
			&PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{
						ID:       1,
						OwnerID:  common.PLAYER_OWNER_ID,
						Category: "M",
						Growth:   0,
						MaxUnits: 200,
						Units:    30,
					},
				},
			},
			[]dto.StatusFleet{},
			30,
			common.PLAYER_OWNER_ID,
			0,
			0,
		},
		// CASE 4: not growing more than MaxUnits
		{
			&PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{
						ID:       1,
						OwnerID:  common.PLAYER_OWNER_ID,
						Category: "M",
						Growth:   5,
						MaxUnits: 200,
						Units:    198,
					},
				},
			},
			[]dto.StatusFleet{},
			200,
			common.PLAYER_OWNER_ID,
			0,
			0,
		},
		// CASE 5: the player attacks a neutral planet.
		// The units do not grow.
		// The planet remains neutral, even at 0 units.
		{
			&PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{
						ID:       1,
						OwnerID:  common.NEUTRAL_OWNER_ID,
						Category: "M",
						Growth:   5,
						MaxUnits: 40,
						Units:    30,
					},
				},
			},
			[]dto.StatusFleet{
				{OwnerID: common.PLAYER_OWNER_ID, Units: 30},
			},
			0,
			common.NEUTRAL_OWNER_ID,
			30,
			0,
		},
		// CASE 6: the player attacks an enemy planet (OwnerID > 1).
		// The units grow before the fleets attacks.
		// The planet remains owned by the enemy.
		{
			&PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{
						ID:       1,
						OwnerID:  2,
						Category: "M",
						Growth:   5,
						MaxUnits: 40,
						Units:    30,
					},
				},
			},
			[]dto.StatusFleet{
				{OwnerID: 1, Units: 31},
			},
			4,
			2,
			31,
			0,
		},
		// CASE 7: the player earns an enemy planet (OwnerID > 1).
		{
			&PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{
						ID:       1,
						OwnerID:  2,
						Category: "M",
						Growth:   5,
						MaxUnits: 40,
						Units:    30,
					},
				},
			},
			[]dto.StatusFleet{
				{OwnerID: 1, Units: 60},
			},
			25,
			common.PLAYER_OWNER_ID,
			60,
			0,
		},
	}

	for index, testCase := range testCases {
		// Act
		testCase.Input.NextTurn(testCase.InputFleets)

		// Assert
		units := testCase.Input.Turns[1].Units
		ownerID := testCase.Input.Turns[1].OwnerID
		totalUnitsSent := testCase.Input.TotalUnitsSent
		totalEnemyUnitsSent := testCase.Input.TotalEnemyUnitsSent

		if testCase.ExpectedUnits != units {
			t.Errorf("TestNextTurn(%d): expected units %s, actual %s", index,
				testCase.ExpectedUnits, units)
		}

		if testCase.ExpectedOwnerID != ownerID {
			t.Errorf("TestNextTurn(%d): expected owner %s, actual %s", index,
				testCase.ExpectedOwnerID, ownerID)
		}

		if testCase.ExpectedTotalUnitsSent != totalUnitsSent {
			t.Errorf("TestNextTurn(%d): expected totalUnitsSent %s, actual %s", index,
				testCase.ExpectedTotalUnitsSent, totalUnitsSent)
		}

		if testCase.ExpectedTotalEnemyUnitsSent != totalEnemyUnitsSent {
			t.Errorf("TestNextTurn(%d): expected totalEnemyUnitsSent %s, actual %s", index,
				testCase.ExpectedTotalUnitsSent, totalEnemyUnitsSent)
		}
	}
}

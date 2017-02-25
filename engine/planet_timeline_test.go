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

type ExpectedResultNextTurn struct {
	Units               int16
	OwnerID             int16
	TotalUnitsSent      int16
	TotalEnemyUnitsSent int16
}

func TestNextTurn(t *testing.T) {
	// Arrange
	var testCases = []struct {
		Input       *PlanetTimeline
		InputFleets []dto.StatusFleet
		Expected    ExpectedResultNextTurn
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
			ExpectedResultNextTurn{
				Units:               35,
				OwnerID:             common.PLAYER_OWNER_ID,
				TotalUnitsSent:      0,
				TotalEnemyUnitsSent: 0,
			},
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
			ExpectedResultNextTurn{
				Units:               30,
				OwnerID:             common.NEUTRAL_OWNER_ID,
				TotalUnitsSent:      0,
				TotalEnemyUnitsSent: 0,
			},
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
			ExpectedResultNextTurn{
				Units:               30,
				OwnerID:             common.PLAYER_OWNER_ID,
				TotalUnitsSent:      0,
				TotalEnemyUnitsSent: 0,
			},
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
			ExpectedResultNextTurn{
				Units:               200,
				OwnerID:             common.PLAYER_OWNER_ID,
				TotalUnitsSent:      0,
				TotalEnemyUnitsSent: 0,
			},
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
			ExpectedResultNextTurn{
				Units:               0,
				OwnerID:             common.NEUTRAL_OWNER_ID,
				TotalUnitsSent:      30,
				TotalEnemyUnitsSent: 0,
			},
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
			ExpectedResultNextTurn{
				Units:               4,
				OwnerID:             2,
				TotalUnitsSent:      31,
				TotalEnemyUnitsSent: 0,
			},
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
			ExpectedResultNextTurn{
				Units:               25,
				OwnerID:             common.PLAYER_OWNER_ID,
				TotalUnitsSent:      60,
				TotalEnemyUnitsSent: 0,
			},
		},
		// CASE 8: the player sends units to its own planet
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
			[]dto.StatusFleet{
				{OwnerID: 1, Units: 60},
			},
			ExpectedResultNextTurn{
				Units:               95,
				OwnerID:             common.PLAYER_OWNER_ID,
				TotalUnitsSent:      60,
				TotalEnemyUnitsSent: 0,
			},
		},
		// CASE 9: the enemy takes the planet
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
			[]dto.StatusFleet{
				{OwnerID: 2, Units: 60},
			},
			ExpectedResultNextTurn{
				Units:               25,
				OwnerID:             2,
				TotalUnitsSent:      0,
				TotalEnemyUnitsSent: 60,
			},
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

		// I do that because the index of the comments starts at 1 '-__-
		index += 1

		if testCase.Expected.Units != units {
			t.Errorf("TestNextTurn(%d): expected units %s, actual %s", index,
				testCase.Expected.Units, units)
		}

		if testCase.Expected.OwnerID != ownerID {
			t.Errorf("TestNextTurn(%d): expected owner %s, actual %s", index,
				testCase.Expected.OwnerID, ownerID)
		}

		if testCase.Expected.TotalUnitsSent != totalUnitsSent {
			t.Errorf("TestNextTurn(%d): expected totalUnitsSent %s, actual %s", index,
				testCase.Expected.TotalUnitsSent, totalUnitsSent)
		}

		if testCase.Expected.TotalEnemyUnitsSent != totalEnemyUnitsSent {
			t.Errorf("TestNextTurn(%d): expected totalEnemyUnitsSent %s, actual %s", index,
				testCase.Expected.TotalEnemyUnitsSent, totalEnemyUnitsSent)
		}
	}
}

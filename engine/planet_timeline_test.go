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
		Input           *PlanetTimeline
		InputFleets     []dto.StatusFleet
		ExpectedUnits   int16
		ExpectedOwnerID uint16
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
				{OwnerID: 1, Units: 30},
			},
			0,
			common.NEUTRAL_OWNER_ID,
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
		},
	}

	for index, testCase := range testCases {
		// Act
		testCase.Input.NextTurn(testCase.InputFleets)

		// Assert
		units := testCase.Input.Turns[1].Units
		ownerID := testCase.Input.Turns[1].OwnerID

		if testCase.ExpectedUnits != units {
			t.Errorf("TestInitDistanceMap(%d): expected units %s, actual %s", index,
				testCase.ExpectedUnits, units)
		}

		if testCase.ExpectedOwnerID != ownerID {
			t.Errorf("TestInitDistanceMap(%d): expected owner %s, actual %s", index,
				testCase.ExpectedOwnerID, ownerID)
		}
	}
}

func TestComputeOwnership(t *testing.T) {
	// Arrange
	testCases := []struct {
		Input    *PlanetTimeline
		Expected int
	}{
		// Case 1: the player has lost the ownership of the planet -> ownership -50
		{
			&PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 100, MaxUnits: 200},
					{ID: 1, OwnerID: 2, Units: 50, MaxUnits: 200},
				},
			},
			-50,
		},
		// Case 2: the player keeps the ownership -> ownership 150
		{
			&PlanetTimeline{
				Turns: []dto.StatusPlanet{
					{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 100, MaxUnits: 200},
					{ID: 1, OwnerID: common.PLAYER_OWNER_ID, Units: 150, MaxUnits: 200},
				},
			},
			150,
		},
		// Case 3: nothing to compute
		{
			&PlanetTimeline{
				Turns: []dto.StatusPlanet{},
			},
			0,
		},
	}

	for index, testCase := range testCases {
		// Arrange
		planetTimeline := testCase.Input

		// Act
		actual := planetTimeline.Ownership()

		// Assert
		if actual != testCase.Expected {
			t.Errorf("TestComputeOwnership(%d): expected (%d), actual was (%d)", index, testCase.Expected, actual)
		}
	}
}

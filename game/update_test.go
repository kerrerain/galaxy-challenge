package game

import (
	"github.com/magleff/galaxy-challenge/dto"
	"testing"
)

func TestUpdateMap(t *testing.T) {
	// Arrange
	gameMap := &Map{}

	status := dto.Status{
		Planets: []dto.StatusPlanet{
			{}, {},
		},
		Fleets: []dto.StatusFleet{
			{}, {}, {},
		},
	}

	// Act
	gameMap.Update(status)

	// Assert
	if len(gameMap.Planets) != len(status.Planets) {
		t.Errorf("There should be %d planets after update.", len(status.Planets))
	}

	if len(gameMap.Fleets) != len(status.Fleets) {
		t.Errorf("There should be %d planets after update.", len(status.Fleets))
	}
}

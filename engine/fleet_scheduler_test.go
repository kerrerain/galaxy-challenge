package engine

import (
	"github.com/magleff/galaxy-challenge/dto"
	"testing"
)

func TestCreateFleetScheduler(t *testing.T) {
	// Arrange
	fleets := []dto.StatusFleet{
		{TargetID: 1, Left: 4},
		{TargetID: 2, Left: 1},
		{TargetID: 1, Left: 4},
		{TargetID: 2, Left: 4},
	}

	// Act
	fleetScheduler := CreateFleetScheduler(fleets)

	// Assert
	if len(fleetScheduler.TurnFleetsForPlanet(4, 1)) != 2 {
		t.Error("TestCreateFleetScheduler: should init the fleets.")
	}
}

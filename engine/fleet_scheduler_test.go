package engine

import (
	"github.com/magleff/galaxy-challenge/dto"
	"reflect"
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

func TestFleets(t *testing.T) {
	// Arrange
	fleets := []dto.StatusFleet{
		{OwnerID: 1, Left: 3, TargetID: 1, Units: 50},
	}
	fleetScheduler := CreateFleetScheduler(fleets)

	// Act
	actual := fleetScheduler.Fleets()

	// Assert
	if !reflect.DeepEqual(fleets, actual) {
		t.Errorf("TestFleets: expected %v, was %v", fleets, actual)
	}
}

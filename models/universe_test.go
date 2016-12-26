package models

import (
	"reflect"
	"testing"
)

var universe Universe

func TestUpdate(t *testing.T) {
	// Arrange
	var testCases = []struct {
		Input    Request
		Expected Universe
	}{
		{
			Request{
				Planets: []Planet{
					{ID: 1, OwnerID: 2}, {ID: 2, OwnerID: 2}, {ID: 3, OwnerID: 1},
				},
				Fleets: []Fleet{
					{OwnerID: 1}, {OwnerID: 1},
				},
			},
			Universe{
				Planets: map[uint16][]Planet{
					1: []Planet{
						{ID: 3, OwnerID: 1},
					},
					2: []Planet{
						{ID: 1, OwnerID: 2}, {ID: 2, OwnerID: 2},
					},
				},
				Fleets: map[uint16][]Fleet{
					1: []Fleet{
						{OwnerID: 1}, {OwnerID: 1},
					},
				},
				Stargates: map[uint16][]*Stargate{},
			},
		},
	}

	for index, testCase := range testCases {
		// Arrange
		actual := CreateNewUniverse()

		// Act
		actual.Update(testCase.Input)

		// Assert
		if !reflect.DeepEqual(actual, testCase.Expected) {
			t.Errorf("TestUpdate(%d): expected %s, actual %s", index, testCase.Expected, actual)
		}
	}
}

func TestOpenStargate(t *testing.T) {
	// Arrange
	var testCases = []struct {
		SourceID uint16
		TargetID uint16
		Expected Universe
	}{
		{
			1,
			2,
			Universe{
				Stargates: map[uint16][]*Stargate{
					1: []*Stargate{
						{
							SourcePlanetID: 1,
							TargetPlanetID: 2,
						},
					},
					2: []*Stargate{
						{
							SourcePlanetID: 1,
							TargetPlanetID: 2,
						},
					},
				},
			},
		},
	}

	for index, testCase := range testCases {
		// Arrange
		actual := CreateNewUniverse()

		// Act
		actual.OpenStargate(testCase.SourceID, testCase.TargetID)

		// Assert
		if !reflect.DeepEqual(actual, testCase.Expected) {
			t.Errorf("TestOpenStargate(%d): expected %s, actual %s", index, testCase.Expected, actual)
		}
	}
}

package game

import (
	"testing"
)

func TestDistanceFromTo(t *testing.T) {
	// Arrange
	var testCases = []struct {
		InputGame     *Game
		InputSourceID uint16
		InputTargetID uint16
		Expected      float64
	}{
		{
			&Game{
				PlanetsByID: map[uint16]*Planet{
					1: &Planet{
						Distances: map[uint16]*Distance{
							2: &Distance{Distance: 40},
						},
					},
				},
			},
			1,
			2,
			40,
		},
	}

	for index, testCase := range testCases {
		// Arrange
		input := testCase.InputGame

		// Act
		actual := input.DistanceFromTo(testCase.InputSourceID, testCase.InputTargetID)

		// Assert
		if actual != testCase.Expected {
			t.Errorf("TestRun(%d): expected %s, actual %s", index, testCase.Expected, actual)
		}
	}
}

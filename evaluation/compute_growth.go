package evaluation

import (
	"github.com/magleff/galaxy-challenge/engine"
	"math"
)

func ComputeGrowth(planet *engine.PlanetTimeline) int16 {
	growth := int16(0)
	lastTurn := int16(0)

	if len(planet.Turns) > 1 {
		lastTurn = planet.Turns[0].Units
	} else {
		return growth
	}

	for _, turn := range planet.Turns[1:len(planet.Turns)] {
		growth += turn.Units - lastTurn
		lastTurn = turn.Units
	}

	return int16(math.Ceil(float64(growth / int16(len(planet.Turns)-1))))
}

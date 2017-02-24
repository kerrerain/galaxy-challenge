package simulation

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/game"
)

func ComputeTaken(gameMap *game.Map, targetID int16) bool {
	timeline := engine.CreateTimelineForPlanets(gameMap, []int16{targetID})

	for i := 0; i < common.SIMULATION_HORIZON; i++ {
		timeline.NextTurn()
	}

	return timeline.PlanetTimelinesMap[targetID].CurrentTurn().OwnerID == common.PLAYER_OWNER_ID
}

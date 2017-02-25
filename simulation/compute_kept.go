package simulation

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/game"
)

func ComputeKept(gameMap *game.Map, targetID int16, move dto.Move) bool {
	timeline := engine.CreateTimelineForPlanets(gameMap, []int16{targetID})
	timeline.ScheduleMoveForNextTurn(common.PLAYER_OWNER_ID, move)

	for i := 0; i < common.SIMULATION_HORIZON; i++ {
		timeline.NextTurn()
	}

	return timeline.PlanetTimelinesMap[targetID].CurrentTurn().OwnerID == common.PLAYER_OWNER_ID
}

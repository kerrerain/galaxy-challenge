package simulation

import (
	"github.com/magleff/galaxy-challenge/command"
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/game"
)

func ComputeKept(gameMap *game.Map, sourceID int16, order command.Order) bool {
	timeline := engine.CreateTimelineForPlanets(gameMap, []int16{sourceID})

	commander := command.CreateCommanderFromTimeline(timeline)
	commander.SendOrder(order)

	timeline.ScheduleMoveForNextTurn(common.PLAYER_OWNER_ID, commander.BuildMove())

	for i := 0; i < common.SIMULATION_HORIZON; i++ {
		timeline.NextTurn()
	}

	return timeline.PlanetTimelinesMap[sourceID].CurrentTurn().OwnerID == common.PLAYER_OWNER_ID
}

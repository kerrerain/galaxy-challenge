package simulation

import (
	"github.com/magleff/galaxy-challenge/command"
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/game"
)

func ComputeKept(gameMap *game.Map, sourceID int16, targetUnits int16) bool {
	timeline := engine.CreateTimelineForPlanets(gameMap, []int16{sourceID})

	for i := 0; i < common.SIMULATION_HORIZON; i++ {
		commander := command.CreateCommanderFromTimeline(timeline)
		commander.SendOrder(command.Order{
			SourceID: sourceID,
			TargetID: 0,
			Units:    targetUnits + 1,
		})

		timeline.ScheduleMoveForNextTurn(common.PLAYER_OWNER_ID, commander.BuildMove())
		timeline.NextTurn()
	}

	return timeline.PlanetTimelinesMap[sourceID].CurrentTurn().OwnerID == common.PLAYER_OWNER_ID
}

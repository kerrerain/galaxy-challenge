package leraje

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/evaluation"
	"github.com/magleff/galaxy-challenge/game"
)

// Spirit of strength and humility.
//
// An IA quickly simulating the future of the game, in order to avoid dumb moves.
// Hopefully better than Paimon.
//
func Run(gameMap *game.Map) dto.Move {
	timeline := runBasicTimeline(gameMap)
	analysis := evaluation.Run(timeline)

	return dto.Move{}
}

func runBasicTimeline(gameMap *game.Map) engine.Timeline {
	timeline := engine.CreateTimeline(gameMap)

	for i := 0; i < common.SIMULATION_HORIZON; i++ {
		timeline.NextTurn()
	}

	return timeline
}

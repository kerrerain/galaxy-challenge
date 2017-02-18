package leraje

import (
	"github.com/magleff/galaxy-challenge/command"
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/evaluation"
	"github.com/magleff/galaxy-challenge/game"
	"sort"
)

// Spirit of strength and humility.
//
// An IA quickly simulating the future of the game, in order to avoid dumb moves.
// Hopefully better than Paimon.
//
func Run(gameMap *game.Map) dto.Move {
	commander := command.CreateCommander(gameMap)
	timeline := runBasicTimeline(gameMap)
	result := evaluation.Run(timeline)
	sort.Sort(evaluation.ByLowestDistanceToPlayer(result.Planets))

	ownPlanets := dto.FilterStatusPlanets(gameMap.Planets, func(planet dto.StatusPlanet) bool {
		return planet.OwnerID == common.PLAYER_OWNER_ID
	})

	for _, planet := range ownPlanets {
		for _, target := range result.Planets {
			if target.Origin.OwnerID != common.PLAYER_OWNER_ID && target.Loss < 1 {
				commander.SendOrder(command.Order{
					SourceID: planet.ID,
					TargetID: target.ID,
					Units:    uint16(target.Origin.Units) + 1,
				})
			}
		}
	}

	return commander.BuildMove()
}

func runBasicTimeline(gameMap *game.Map) engine.Timeline {
	timeline := engine.CreateTimeline(gameMap)

	for i := 0; i < common.SIMULATION_HORIZON; i++ {
		timeline.NextTurn()
	}

	return timeline
}

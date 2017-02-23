package simulation

import (
	"github.com/magleff/galaxy-challenge/command"
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/game"
)

func ComputeFullInvasionCost(gameMap *game.Map, targetIDs []int16, sourceIDs []int16, enemySourceIDs []int16) []InvasionCost {
	invasionCost := make([]InvasionCost, len(targetIDs))

	for i, targetID := range targetIDs {
		invasionCost[i] = InvasionCost{
			ID:   targetID,
			Cost: computeFullPlanetInvasionCost(gameMap, targetID, sourceIDs, enemySourceIDs),
		}
	}

	return invasionCost
}

func computeFullPlanetInvasionCost(gameMap *game.Map, targetID int16, sourceIDs []int16, enemySourceIDs []int16) int16 {
	timeline := engine.CreateTimeline(gameMap)

	for i := 0; i < common.SIMULATION_HORIZON; i++ {
		//currentTurn := timeline.PlanetTimelinesMap[targetID].CurrentTurn()

		/*		if currentTurn.OwnerID != common.PLAYER_OWNER_ID {


				}*/
		commander := command.CreateCommanderFromTimeline(timeline)
		commander.AllOutAttackOnPlanet(targetID, sourceIDs)

		enemyCommander := command.CreateCommanderFromTimeline(timeline)
		enemyCommander.AllOutAttackOnPlanet(targetID, enemySourceIDs)

		timeline.ScheduleMoveForNextTurn(common.PLAYER_OWNER_ID, commander.BuildMove())
		timeline.ScheduleMoveForNextTurn(2, enemyCommander.BuildMove())
		timeline.NextTurn()
	}

	return timeline.PlanetTimelinesMap[targetID].TotalUnitsSent - timeline.PlanetTimelinesMap[targetID].TotalEnemyUnitsSent
}

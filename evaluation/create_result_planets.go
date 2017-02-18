package evaluation

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/game"
)

func createResultPlanets(timeline engine.Timeline) []*ResultPlanet {
	resultPlanets := make([]*ResultPlanet, len(timeline.PlanetTimelines))

	for i, planetTimeline := range timeline.PlanetTimelines {
		resultPlanets[i] = &ResultPlanet{
			ID:               planetTimeline.ID,
			Origin:           planetTimeline.Origin,
			Loss:             ComputeLoss(planetTimeline),
			Growth:           ComputeGrowth(planetTimeline),
			DistanceToEnemy:  ComputeMedianDistanceToEnemy(timeline.GameMap, planetTimeline, timeline.PlanetTimelines),
			DistanceToPlayer: ComputeMedianDistanceToPlayer(timeline.GameMap, planetTimeline, timeline.PlanetTimelines),
		}
	}

	return resultPlanets
}

func ComputeMedianDistanceToEnemy(gameMap *game.Map, source *engine.PlanetTimeline,
	targets []*engine.PlanetTimeline) int16 {

	enemiesID := make([]int16, 0)

	for _, target := range targets {
		lastTurn := target.Turns[len(target.Turns)-1]
		if lastTurn.OwnerID != common.PLAYER_OWNER_ID && lastTurn.OwnerID != common.NEUTRAL_OWNER_ID {
			enemiesID = append(enemiesID, target.ID)
		}
	}

	return gameMap.ComputeMedianDistance(source.ID, enemiesID).Turns
}

func ComputeMedianDistanceToPlayer(gameMap *game.Map, source *engine.PlanetTimeline,
	targets []*engine.PlanetTimeline) int16 {

	playersID := make([]int16, 0)

	for _, target := range targets {
		lastTurn := target.Turns[len(target.Turns)-1]
		if lastTurn.OwnerID == common.PLAYER_OWNER_ID {
			playersID = append(playersID, target.ID)
		}
	}

	return gameMap.ComputeMedianDistance(source.ID, playersID).Turns
}

package simulation

import (
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
)

func ComputeROI(gameMap *game.Map, sourceID int16, targetID int16, playerID int16) float64 {
	sourcePlanet := dto.GetByID(gameMap.Planets, sourceID)
	targetPlanet := dto.GetByID(gameMap.Planets, targetID)

	if targetPlanet.Growth == 0 {
		return 0
	} else {
		return float64((targetPlanet.Units / targetPlanet.Growth) +
			(targetPlanet.Units / sourcePlanet.Growth) +
			2*gameMap.DistanceMap[sourceID][targetID].Turns)
	}
}

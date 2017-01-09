package paimon

import (
	"github.com/magleff/galaxy-challenge/dto"
	"github.com/magleff/galaxy-challenge/game"
)

// Spirit of maniacal love and chaos.
//
// A simple minded IA, that will try to invade the nearest and best planets,
// without analyzing accurately the consequences.
//
func ComputeMove(gameMap *game.Map) dto.Move {
	planetEvaluations := evaluate(gameMap)
	commander := CreateCommander(gameMap)

	return commander.PlanInvasion(planetEvaluations)
}

package paimon

import (
	"github.com/magleff/galaxy-challenge/dto"
	//"github.com/magleff/galaxy-challenge/engine"
	"github.com/magleff/galaxy-challenge/game"
)

// Spirit of maniacal love and chaos.
//
// A simple minded IA, that will try to invade the nearest and best planets,
// without analyzing accurately the consequences.
//
func Run(gameMap *game.Map) dto.Move {
	//timeline := engine.CreateTimeline(gameMap)
	evaluation := CreateEvaluation(gameMap)
	commander := CreateCommander(gameMap)

	return commander.PlanInvasion(evaluation.Run())
}

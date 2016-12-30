package core

import (
	"github.com/magleff/galaxy-challenge/models/game"
	"github.com/magleff/galaxy-challenge/models/move"
	"github.com/magleff/galaxy-challenge/simulation"
)

func MakeMove(g *game.Game) (*move.Move, error) {
	return simulation.SimulateSimple(g), nil
}

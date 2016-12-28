package core

import (
	"github.com/magleff/galaxy-challenge/mappers"
	"github.com/magleff/galaxy-challenge/models/game"
	"github.com/magleff/galaxy-challenge/models/status"
)

func UpdateGame(status status.Game, g *game.Game) {
	if !g.Initialized {
		initializeGame(status, g)
	} else {
		updatePlanets(status.Planets, g)
	}

	g.Turn = status.Config.Turn
	g.MaxTurn = status.Config.MaxTurn
	g.Fleets = status.Fleets
}

func initializeGame(status status.Game, g *game.Game) {
	initializePlanets(status.Planets, g)

	g.InitializeDistances()
	g.Initialized = true
}

func initializePlanets(planets []status.Planet, g *game.Game) {
	for _, planet := range planets {
		g.InitPlanet(mappers.ToGamePlanet(planet))
	}
}

func updatePlanets(planets []status.Planet, g *game.Game) {
	for _, planet := range planets {
		mappers.UpdateGamePlanet(g.PlanetsByID[planet.ID], planet)
	}
}

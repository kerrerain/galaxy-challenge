package core

import (
	"github.com/magleff/galaxy-challenge/globals"
	"github.com/magleff/galaxy-challenge/models/game"
	"github.com/magleff/galaxy-challenge/models/move"
)

func MakeMove(g *game.Game) (move.Move, error) {
	response := move.CreateMove()

	for _, planet := range g.Planets {
		if planet.OwnerID == globals.PLAYER_OWNER_ID {
			response.Fleets = sendTestFleet(planet)
		}
	}

	return response, nil
}

func sendTestFleet(planet *game.Planet) []move.Fleet {
	fleets := make([]move.Fleet, 0)
	minDist := planet.DistancesSorted[0]

	if planet.Units > planet.MaxUnits/2 {
		fleets = append(fleets, move.Fleet{
			SourceID: planet.ID,
			TargetID: minDist.Planet.ID,
			Units:    planet.Units / 3,
		})
	}

	return fleets
}

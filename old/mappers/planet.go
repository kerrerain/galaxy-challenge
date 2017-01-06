package mappers

import (
	"github.com/magleff/galaxy-challenge/models/game"
	"github.com/magleff/galaxy-challenge/models/status"
)

func ToGamePlanet(planet status.Planet) *game.Planet {
	return &game.Planet{
		ID:        planet.ID,
		X:         planet.X,
		Y:         planet.Y,
		OwnerID:   planet.OwnerID,
		Units:     planet.Units,
		MaxUnits:  planet.MaxUnits,
		Growth:    planet.Growth,
		Category:  planet.Category,
		Distances: make(map[uint16]*game.Distance),
	}
}

func UpdateGamePlanet(planetInGame *game.Planet, planet status.Planet) {
	planetInGame.ID = planet.ID
	planetInGame.X = planet.X
	planetInGame.Y = planet.Y
	planetInGame.OwnerID = planet.OwnerID
	planetInGame.Units = planet.Units
	planetInGame.MaxUnits = planet.MaxUnits
	planetInGame.Growth = planet.Growth
	planetInGame.Category = planet.Category
}

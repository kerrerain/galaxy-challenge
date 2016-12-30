package game

import (
	"github.com/magleff/galaxy-challenge/models/status"
)

type Game struct {
	Planets          []*Planet
	PlanetsByID      map[uint16]*Planet
	PlanetsByOwnerID map[uint16][]*Planet
	Fleets           []status.Fleet
	Turn             uint16
	MaxTurn          uint16
	Initialized      bool
}

func (g *Game) InitPlanet(planet *Planet) {
	g.Planets = append(g.Planets, planet)
	g.PlanetsByID[planet.ID] = planet
}

func (g *Game) UpdatePlanetsByOwnerID() {
	g.PlanetsByOwnerID = make(map[uint16][]*Planet)

	for _, planet := range g.Planets {
		if g.PlanetsByOwnerID[planet.OwnerID] == nil {
			g.PlanetsByOwnerID[planet.OwnerID] = make([]*Planet, 0)
		}
		g.PlanetsByOwnerID[planet.OwnerID] = append(g.PlanetsByOwnerID[planet.OwnerID], planet)
	}
}

func (g *Game) InitializeDistances() {
	for _, planet := range g.Planets {
		planet.InitializeDistances(g.Planets)
	}
}

func (g Game) DistanceFromTo(sourceID uint16, targetID uint16) float64 {
	return g.PlanetsByID[sourceID].Distances[targetID].Distance
}

func (g Game) CopyPlanets() []*Planet {
	planets := make([]*Planet, 0)

	for _, planet := range g.Planets {
		planets = append(planets, planet.Copy())
	}

	return planets
}

func CreateNewGame() *Game {
	return &Game{
		Planets:          make([]*Planet, 0),
		PlanetsByID:      make(map[uint16]*Planet),
		PlanetsByOwnerID: make(map[uint16][]*Planet),
	}
}

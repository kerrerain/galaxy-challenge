package game

import (
	"github.com/magleff/galaxy-challenge/models/status"
)

type Game struct {
	Planets     []*Planet
	PlanetsByID map[uint16]*Planet
	Fleets      []status.Fleet
	Turn        uint16
	MaxTurn     uint16
	Initialized bool
}

func (g *Game) InitPlanet(planet *Planet) {
	g.Planets = append(g.Planets, planet)
	g.PlanetsByID[planet.ID] = planet
}

func (g *Game) InitializeDistances() {
	for _, planet := range g.Planets {
		planet.InitializeDistances(g.Planets)
	}
}

func CreateNewGame() *Game {
	return &Game{
		Planets:     make([]*Planet, 0),
		PlanetsByID: make(map[uint16]*Planet),
	}
}

/*func (u *Universe) Update(request Request) {
	u.UpdatePlanets(request.Planets)
	u.UpdateFleets(request.Fleets)
}

func (u *Universe) UpdatePlanets(planets []Planet) {
	u.Planets = make(map[uint16][]Planet)

	for _, planet := range planets {
		id := planet.OwnerID

		if u.Planets[id] == nil {
			u.Planets[id] = make([]Planet, 0)
		}

		u.Planets[id] = append(u.Planets[id], planet)
	}
}

func (u *Universe) UpdateFleets(fleets []Fleet) {
	u.Fleets = make(map[uint16][]Fleet)

	for _, fleet := range fleets {
		id := fleet.OwnerID

		if u.Fleets[id] == nil {
			u.Fleets[id] = make([]Fleet, 0)
		}

		u.Fleets[id] = append(u.Fleets[id], fleet)
	}
}

func (u *Universe) OpenStargate(sourcePlanetID uint16, targetPlanetID uint16) {
	log.Println("Opening a stargate from", sourcePlanetID, "to", targetPlanetID)

	if u.Stargates[sourcePlanetID] == nil {
		u.Stargates[sourcePlanetID] = make([]*Stargate, 0)
	}

	stargate := &Stargate{
		SourcePlanetID: sourcePlanetID,
		TargetPlanetID: targetPlanetID,
	}

	u.Stargates[sourcePlanetID] = append(u.Stargates[sourcePlanetID], stargate)
}*/

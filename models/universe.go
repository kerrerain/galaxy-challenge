package models

import (
	"log"
)

type Universe struct {
	Planets   map[uint16][]Planet
	Fleets    map[uint16][]Fleet
	Stargates map[uint16][]*Stargate
}

func (u *Universe) Update(request Request) {
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
}

func CreateNewUniverse() Universe {
	return Universe{
		Stargates: make(map[uint16][]*Stargate),
	}
}

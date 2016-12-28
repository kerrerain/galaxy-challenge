package game

import (
	"math"
	"sort"
)

type Planet struct {
	ID              uint16
	X               float64
	Y               float64
	OwnerID         uint16
	Units           uint16
	MaxUnits        uint16
	Growth          uint16
	Category        string
	Distances       map[uint16]*Distance
	DistancesSorted []*Distance
}

func (p Planet) Habitable() bool {
	return p.Category != "D" &&
		p.Category != "J" &&
		p.Category != "N"
}

func (p Planet) ComputeDistanceToPlanet(p2 *Planet) float64 {
	return math.Sqrt(math.Pow(p2.X-p.X, 2) + math.Pow(p2.Y-p.Y, 2))
}

func (p *Planet) InitializeDistances(planets []*Planet) {
	distancesToSort := make([]*Distance, 0)

	for _, planet := range planets {
		if planet.ID != p.ID {
			p.Distances[planet.ID] = &Distance{
				Planet:   planet,
				Distance: p.ComputeDistanceToPlanet(planet),
			}
			distancesToSort = append(distancesToSort, p.Distances[planet.ID])
		}
	}

	// Asc
	sort.Sort(ByDistance(distancesToSort))
	p.DistancesSorted = distancesToSort
}

func FilterPlanets(toFilter []Planet, predicate func(Planet) bool) []Planet {
	filtered := make([]Planet, 0)
	for _, planet := range toFilter {
		if predicate(planet) {
			filtered = append(filtered, planet)
		}
	}
	return filtered
}

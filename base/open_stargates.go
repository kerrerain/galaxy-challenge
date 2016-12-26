package base

import (
	"github.com/magleff/galaxy-challenge/globals"
	"github.com/magleff/galaxy-challenge/models"
	"sort"
)

type Distance struct {
	PlanetID uint16
	Distance float64
}

type ByDistance []Distance

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }

func openStargatesForPlanet(planet models.Planet, universe *models.Universe) {
	distances := make([]Distance, 0)
	targetPlanetIDs := make([]uint16, 0)

	for _, neutralPlanet := range universe.Planets[globals.NEUTRAL_OWNER_ID] {
		distances = append(distances, Distance{
			PlanetID: neutralPlanet.ID,
			Distance: computeDistance(planet, neutralPlanet),
		})
	}

	// Asc
	sort.Sort(ByDistance(distances))

	for _, distance := range distances {
		if len(targetPlanetIDs) < globals.MAX_STARGATES_BY_PLANET {
			targetPlanetIDs = append(targetPlanetIDs, distance.PlanetID)
		}
	}

	for _, targetPlanetID := range targetPlanetIDs {
		universe.OpenStargate(planet.ID, targetPlanetID)
	}
}

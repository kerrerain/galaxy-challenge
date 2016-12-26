package base

import (
	"github.com/magleff/galaxy-challenge/globals"
	"github.com/magleff/galaxy-challenge/models"
)

func Analyse(game models.Game) (models.Response, error) {
	response := models.CreateResponse()

	for _, planet := range game.Universe.Planets[globals.PLAYER_OWNER_ID] {
		if game.Universe.Stargates[planet.ID] == nil {
			openStargatesForPlanet(planet, &game.Universe)
		}
	}

	for _, planet := range game.Universe.Planets[globals.PLAYER_OWNER_ID] {
		for _, stargate := range game.Universe.Stargates[planet.ID] {
			//TODO find target planet Units
			if stargate.SourcePlanetID == planet.ID && planet.Units > planet.MaxUnits/2 {
				response.EngagingFleets = append(response.EngagingFleets, models.EngagingFleet{
					SourceID: stargate.SourcePlanetID,
					TargetID: stargate.TargetPlanetID,
					Units:    planet.Units / globals.MAX_STARGATES_BY_PLANET,
				})
			}
		}
	}

	return response, nil
}

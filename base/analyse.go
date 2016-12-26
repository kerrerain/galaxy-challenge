package base

import (
	"github.com/magleff/galaxy-challenge/models"
)

func Analyse(game models.Game) (models.Response, error) {
	response := models.CreateResponse()

	// TEST
	/*	anotherPlanet := game.Universe.OtherPlanets[0]
		ownPlanet := game.Universe.OwnPlanets[0]

		if ownPlanet.Units > anotherPlanet.Units {
			response.EngagingFleets = []models.EngagingFleet{
				{
					SourceID: ownPlanet.ID,
					TargetID: anotherPlanet.ID,
					Units:    anotherPlanet.Units + 1,
				},
			}
		}*/

	return response, nil
}

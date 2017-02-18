package evaluation

import (
	"github.com/magleff/galaxy-challenge/engine"
)

func Run(timeline engine.Timeline) Result {
	planets := createResultPlanets(timeline)

	result := Result{
		Planets:    planets,
		FrontLimit: ComputeFrontLimit(planets, 5),
	}
	return result
}

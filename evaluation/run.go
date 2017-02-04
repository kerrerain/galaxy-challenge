package evaluation

import (
	"github.com/magleff/galaxy-challenge/engine"
)

func Run(timeline engine.Timeline) Result {
	result := Result{
		Planets: createResultPlanets(timeline),
	}
	return result
}

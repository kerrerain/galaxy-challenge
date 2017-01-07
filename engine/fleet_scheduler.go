package engine

import (
	"github.com/magleff/galaxy-challenge/dto"
	"strconv"
)

type FleetScheduler struct {
	Fleets map[string][]dto.StatusFleet
}

func (f FleetScheduler) TurnFleetsForPlanet(turn int, planetID uint16) []dto.StatusFleet {
	return f.Fleets[computeKey(turn, int(planetID))]
}

func CreateFleetScheduler(fleets []dto.StatusFleet) *FleetScheduler {
	return &FleetScheduler{
		Fleets: initFleets(fleets),
	}
}

func initFleets(fleets []dto.StatusFleet) map[string][]dto.StatusFleet {
	turns := make(map[string][]dto.StatusFleet)

	for _, fleet := range fleets {
		addFleet(turns, fleet)
	}

	return turns
}

func addFleet(turns map[string][]dto.StatusFleet, fleet dto.StatusFleet) {
	key := computeKey(int(fleet.Left), int(fleet.TargetID))

	if turns[key] == nil {
		turns[key] = make([]dto.StatusFleet, 0)
	}

	turns[key] = append(turns[key], fleet)
}

func computeKey(turn int, planetID int) string {
	return strconv.Itoa(turn) + "-" + strconv.Itoa(planetID)
}

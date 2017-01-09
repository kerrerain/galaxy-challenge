package engine

import (
	"github.com/magleff/galaxy-challenge/dto"
	"strconv"
)

type FleetScheduler struct {
	FleetArrivals map[string][]dto.StatusFleet
}

func (f FleetScheduler) TurnFleetsForPlanet(turn int, planetID uint16) []dto.StatusFleet {
	return f.FleetArrivals[computeKey(turn, int(planetID))]
}

func (f *FleetScheduler) AddFleets(fleets []dto.StatusFleet) {
	for _, fleet := range fleets {
		addFleetArrival(f.FleetArrivals, fleet)
	}
}

func CreateFleetScheduler(fleets []dto.StatusFleet) *FleetScheduler {
	return &FleetScheduler{
		FleetArrivals: initFleets(fleets),
	}
}

func initFleets(fleets []dto.StatusFleet) map[string][]dto.StatusFleet {
	fleetArrivals := make(map[string][]dto.StatusFleet)

	for _, fleet := range fleets {
		addFleetArrival(fleetArrivals, fleet)
	}

	return fleetArrivals
}

func addFleetArrival(fleetArrivals map[string][]dto.StatusFleet, fleet dto.StatusFleet) {
	key := computeKey(int(fleet.Left), int(fleet.TargetID))

	if fleetArrivals[key] == nil {
		fleetArrivals[key] = make([]dto.StatusFleet, 0)
	}

	fleetArrivals[key] = append(fleetArrivals[key], fleet)
}

func computeKey(turn int, planetID int) string {
	return strconv.Itoa(turn) + "-" + strconv.Itoa(planetID)
}

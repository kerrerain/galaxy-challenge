package engine

import (
	"github.com/magleff/galaxy-challenge/dto"
	"strconv"
	"strings"
)

const FLEET_ARRIVAL_KEY_SEPARATOR = "-"

type FleetScheduler struct {
	FleetArrivals map[string][]dto.StatusFleet
}

func CreateFleetScheduler(fleets []dto.StatusFleet, turn int) *FleetScheduler {
	return &FleetScheduler{
		FleetArrivals: initFleets(fleets, turn),
	}
}

// Returns the fleets landing on a planet at a specific turn (0, 1, etc).
func (f FleetScheduler) TurnFleetsForPlanet(turn int, planetID int16) []dto.StatusFleet {
	return f.FleetArrivals[computeKey(turn, int(planetID))]
}

func (f *FleetScheduler) AddFleets(fleets []dto.StatusFleet, turn int) {
	for _, fleet := range fleets {
		f.AddFleet(fleet, turn)
	}
}

// Schedules a single fleet, using the "Left" parameter to know when it will land on a planet.
func (f *FleetScheduler) AddFleet(fleet dto.StatusFleet, turn int) {
	addFleetArrival(f.FleetArrivals, fleet, turn)
}

func (f FleetScheduler) Fleets(turn int) []dto.StatusFleet {
	fleets := make([]dto.StatusFleet, 0)

	for k, v := range f.FleetArrivals {
		fleets = append(fleets, fleetsForTurn(k, v, turn)...)
	}

	return fleets
}

func fleetsForTurn(key string, value []dto.StatusFleet, turn int) []dto.StatusFleet {
	fleetTurn := computeTurnFromKey(key)
	finalFleets := make([]dto.StatusFleet, 0)

	for _, fleet := range value {
		left := fleetTurn - turn
		if left >= 0 {
			newFleet := fleet.Copy()
			newFleet.Left = int16(left)
			finalFleets = append(finalFleets, newFleet)
		}
	}

	return finalFleets
}

func initFleets(fleets []dto.StatusFleet, turn int) map[string][]dto.StatusFleet {
	fleetArrivals := make(map[string][]dto.StatusFleet)

	for _, fleet := range fleets {
		addFleetArrival(fleetArrivals, fleet, turn)
	}

	return fleetArrivals
}

func addFleetArrival(fleetArrivals map[string][]dto.StatusFleet, fleet dto.StatusFleet, turn int) {
	key := computeKey(turn+int(fleet.Left), int(fleet.TargetID))

	if fleetArrivals[key] == nil {
		fleetArrivals[key] = make([]dto.StatusFleet, 0)
	}

	fleetArrivals[key] = append(fleetArrivals[key], fleet)
}

func computeKey(turn int, planetID int) string {
	return strconv.Itoa(turn) + FLEET_ARRIVAL_KEY_SEPARATOR + strconv.Itoa(planetID)
}

func computeTurnFromKey(key string) int {
	chunks := strings.Split(key, FLEET_ARRIVAL_KEY_SEPARATOR)
	turn, _ := strconv.Atoi(chunks[0])
	return turn
}

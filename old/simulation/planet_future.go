package simulation

import (
	"github.com/magleff/galaxy-challenge/globals"
	"github.com/magleff/galaxy-challenge/models/status"
)

type PlanetFuture struct {
	ID             uint16
	OwnerID        uint16
	Units          int16
	MaxUnits       int16
	Growth         int16
	FleetsArrivals map[int][]status.Fleet
}

func (p *PlanetFuture) FleetsArrivalsForPlanet(moveFleets []status.Fleet, gameFleets []status.Fleet) {
	p.FleetsArrivals = make(map[int][]status.Fleet)
	fleets := append(moveFleets, gameFleets...)

	for _, fleet := range fleets {
		if fleet.TargetID == p.ID {
			time := int(fleet.Left)

			if p.FleetsArrivals[time] == nil {
				p.FleetsArrivals[time] = make([]status.Fleet, 0)
			}
			p.FleetsArrivals[time] = append(p.FleetsArrivals[time], fleet)
		} else if fleet.SourceID == p.ID {
			p.Units = p.Units - fleet.Units
		}
	}
}

func (p *PlanetFuture) Grow() {
	if p.OwnerID != globals.NEUTRAL_OWNER_ID {
		p.Units = Min(p.MaxUnits, p.Units+p.Growth)
	}
}

func (p *PlanetFuture) SimulateFleetsArrivals(time int) {
	fleetsArrivals := p.FleetsArrivals[time]

	if fleetsArrivals != nil {
		for _, arrival := range fleetsArrivals {
			p.Units = p.Units - arrival.Units

			if p.Units < 0 {
				p.OwnerID = arrival.OwnerID // The other player has earned the planet
				p.Units = -1 * p.Units
			}
		}
	}
}

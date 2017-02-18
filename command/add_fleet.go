package command

import (
	"github.com/magleff/galaxy-challenge/dto"
)

func (c *Commander) AddFleet(fleet dto.MoveFleet) {
	var found []*dto.MoveFleet

	for _, existing := range c.Fleets {
		if existing.SourceID == fleet.SourceID && existing.TargetID == fleet.TargetID {
			found = append(found, existing)
		}
	}

	if len(found) == 1 {
		found[0].Units += fleet.Units
	} else {
		c.Fleets = append(c.Fleets, &fleet)
	}
}

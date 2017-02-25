package command

import (
	"github.com/magleff/galaxy-challenge/dto"
)

func (c *Commander) Reset() {
	c.Fleets = make([]*dto.MoveFleet, 0)
}

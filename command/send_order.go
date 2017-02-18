package command

import (
	"github.com/magleff/galaxy-challenge/dto"
)

func (c *Commander) SendOrder(order Order) {
	available := c.AvailableUnitsOnPlanet[order.SourceID]
	toSend := unitsToSend(available, int16(order.Units))

	c.AddFleet(dto.MoveFleet{
		SourceID: order.SourceID,
		TargetID: order.TargetID,
		Units:    toSend,
	})

	c.AvailableUnitsOnPlanet[order.SourceID] = available - toSend
}

func (c Commander) SendMultipleOrders(orders []Order) {
	for _, order := range orders {
		c.SendOrder(order)
	}
}

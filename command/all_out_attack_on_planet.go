package command

func (c *Commander) AllOutAttackOnPlanet(targetID int16, sourceIDs []int16) {
	for _, sourceID := range sourceIDs {
		c.SendOrder(Order{
			SourceID: sourceID,
			TargetID: targetID,
			Units:    c.AvailableUnitsOnPlanet[sourceID] - 1, // All out attack
		})
	}
}

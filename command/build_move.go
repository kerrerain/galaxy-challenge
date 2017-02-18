package command

import (
	"github.com/magleff/galaxy-challenge/dto"
)

func (c Commander) BuildMove() dto.Move {
	return dto.Move{
		Fleets:        buildFleets(c.Fleets),
		Terraformings: make([]dto.MoveTerraforming, 0),
	}
}

func buildFleets(fleets []*dto.MoveFleet) []dto.MoveFleet {
	result := make([]dto.MoveFleet, 0)

	for _, fleet := range fleets {
		if fleet.Units > 0 {
			result = append(result, *fleet)
		}
	}

	return result
}

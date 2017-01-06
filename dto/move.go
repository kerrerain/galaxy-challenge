package dto

type Move struct {
	Fleets        []MoveFleet        `json:"fleets"`
	Terraformings []MoveTerraforming `json:"terraformings"`
}

func CreateMove() *Move {
	return &Move{
		Fleets:        []MoveFleet{},
		Terraformings: []MoveTerraforming{},
	}
}

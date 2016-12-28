package move

type Move struct {
	Fleets        []Fleet        `json:"fleets"`
	Terraformings []Terraforming `json:"terraformings"`
}

func CreateMove() Move {
	return Move{
		Fleets:        []Fleet{},
		Terraformings: []Terraforming{},
	}
}

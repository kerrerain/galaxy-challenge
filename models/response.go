package models

type Response struct {
	EngagingFleets []EngagingFleet `json:"fleets"`
	Terraformings  []Terraforming  `json:"terraformings"`
}

func CreateResponse() Response {
	return Response{
		EngagingFleets: []EngagingFleet{},
		Terraformings:  []Terraforming{},
	}
}

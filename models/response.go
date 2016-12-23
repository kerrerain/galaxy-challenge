package models

type Response struct {
	EngagingFleets []EngagingFleet `json:"fleets,omitempty"`
	Terraformings  []Terraforming  `json:"terraformings,omitempty"`
}

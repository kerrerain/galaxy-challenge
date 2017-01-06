package dto

type Status struct {
	Planets []StatusPlanet      `json:"planets,omitempty"`
	Fleets  []StatusFleet       `json:"fleets,omitempty"`
	Config  StatusConfiguration `json:"config,omitempty"`
}

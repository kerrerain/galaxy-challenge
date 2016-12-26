package models

type Request struct {
	Planets []Planet      `json:"planets,omitempty"`
	Fleets  []Fleet       `json:"fleets,omitempty"`
	Config  Configuration `json:"config,omitempty"`
}

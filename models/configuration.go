package models

type Configuration struct {
	Turn    uint16 `json:"turn,omitempty"`
	MaxTurn uint16 `json:"maxTurn,omitempty"`
}

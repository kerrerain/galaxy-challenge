package dto

type StatusConfiguration struct {
	Turn    uint16 `json:"turn,omitempty"`
	MaxTurn uint16 `json:"maxTurn,omitempty"`
}

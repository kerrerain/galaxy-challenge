package dto

type StatusConfiguration struct {
	Turn    int16 `json:"turn,omitempty"`
	MaxTurn int16 `json:"maxTurn,omitempty"`
}

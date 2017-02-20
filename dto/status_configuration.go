package dto

type StatusConfiguration struct {
	ID      int16 `json:"id,omitempty"`
	Turn    int16 `json:"turn,omitempty"`
	MaxTurn int16 `json:"maxTurn,omitempty"`
}

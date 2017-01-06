package dto

type StatusPlanet struct {
	ID       uint16  `json:"id,omitempty"`
	X        float64 `json:"x,omitempty"`
	Y        float64 `json:"y,omitempty"`
	OwnerID  uint16  `json:"owner,omitempty"`
	Units    int16   `json:"units,omitempty"`
	MaxUnits int16   `json:"mu,omitempty"`
	Growth   int16   `json:"gr,omitempty"`
	Category string  `json:"class,omitempty"`
}

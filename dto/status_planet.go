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

func (p StatusPlanet) Copy() StatusPlanet {
	return StatusPlanet{
		ID:       p.ID,
		X:        p.X,
		Y:        p.Y,
		OwnerID:  p.OwnerID,
		Units:    p.Units,
		MaxUnits: p.MaxUnits,
		Growth:   p.Growth,
		Category: p.Category,
	}
}

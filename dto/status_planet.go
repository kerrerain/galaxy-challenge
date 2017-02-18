package dto

type StatusPlanet struct {
	ID       int16   `json:"id,omitempty"`
	X        float64 `json:"x,omitempty"`
	Y        float64 `json:"y,omitempty"`
	OwnerID  int16   `json:"owner,omitempty"`
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

func FilterStatusPlanets(toFilter []StatusPlanet, predicate func(StatusPlanet) bool) []StatusPlanet {
	filtered := make([]StatusPlanet, 0)
	for _, planet := range toFilter {
		if predicate(planet) {
			filtered = append(filtered, planet)
		}
	}
	return filtered
}

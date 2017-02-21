package dto

type StatusPlanet struct {
	ID       int16   `json:"id"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	OwnerID  int16   `json:"owner"`
	Units    int16   `json:"units"`
	MaxUnits int16   `json:"mu"`
	Growth   int16   `json:"gr"`
	Category string  `json:"class"`
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

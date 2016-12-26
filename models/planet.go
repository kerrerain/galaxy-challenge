package models

type Planet struct {
	ID       uint16  `json:"id,omitempty"`
	X        float32 `json:"x,omitempty"`
	Y        float32 `json:"y,omitempty"`
	OwnerID  uint16  `json:"owner,omitempty"`
	Units    uint16  `json:"units,omitempty"`
	MaxUnits uint16  `json:"mu,omitempty"`
	Growth   uint16  `json:"gr,omitempty"`
	Category string  `json:"class,omitempty"`
}

func (p Planet) Habitable() bool {
	return p.Category != "D" &&
		p.Category != "J" &&
		p.Category != "N"
}

func FilterPlanets(toFilter []Planet, predicate func(Planet) bool) []Planet {
	filtered := make([]Planet, 0)
	for _, planet := range toFilter {
		if predicate(planet) {
			filtered = append(filtered, planet)
		}
	}
	return filtered
}

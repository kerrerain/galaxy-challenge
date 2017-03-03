package dto

type StatusFleet struct {
	OwnerID  int16 `json:"owner"`
	Units    int16 `json:"units"`
	SourceID int16 `json:"from"`
	TargetID int16 `json:"to"`
	Turns    int16 `json:"turns"`
	Left     int16 `json:"left"`
}

func (f StatusFleet) Copy() StatusFleet {
	return StatusFleet{
		OwnerID:  f.OwnerID,
		Units:    f.Units,
		SourceID: f.SourceID,
		TargetID: f.TargetID,
		Turns:    f.Turns,
		Left:     f.Left,
	}
}

func FilterStatusFleets(toFilter []StatusFleet, predicate func(StatusFleet) bool) []StatusFleet {
	filtered := make([]StatusFleet, 0)
	for _, fleet := range toFilter {
		if predicate(fleet) {
			filtered = append(filtered, fleet)
		}
	}
	return filtered
}

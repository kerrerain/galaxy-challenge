package dto

type StatusFleet struct {
	OwnerID  int16 `json:"owner"`
	Units    int16 `json:"units"`
	SourceID int16 `json:"from"`
	TargetID int16 `json:"to"`
	Turns    int16 `json:"turns"`
	Left     int16 `json:"left"`
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

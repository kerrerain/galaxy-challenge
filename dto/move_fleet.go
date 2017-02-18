package dto

type MoveFleet struct {
	SourceID int16 `json:"source,omitempty"`
	TargetID int16 `json:"target,omitempty"`
	Units    int16 `json:"units,omitempty"`
}

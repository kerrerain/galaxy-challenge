package dto

type MoveFleet struct {
	SourceID uint16 `json:"source,omitempty"`
	TargetID uint16 `json:"target,omitempty"`
	Units    int16  `json:"units,omitempty"`
}

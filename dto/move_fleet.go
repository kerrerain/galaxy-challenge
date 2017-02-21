package dto

type MoveFleet struct {
	SourceID int16 `json:"source"`
	TargetID int16 `json:"target"`
	Units    int16 `json:"units"`
}

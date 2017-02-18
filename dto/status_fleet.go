package dto

type StatusFleet struct {
	OwnerID  int16 `json:"owner,omitempty"`
	Units    int16 `json:"units,omitempty"`
	SourceID int16 `json:"from,omitempty"`
	TargetID int16 `json:"to,omitempty"`
	Turns    int16 `json:"turns,omitempty"`
	Left     int16 `json:"left,omitempty"`
}

package dto

type StatusFleet struct {
	OwnerID  uint16 `json:"owner,omitempty"`
	Units    int16  `json:"units,omitempty"`
	SourceID uint16 `json:"from,omitempty"`
	TargetID uint16 `json:"to,omitempty"`
	Turns    uint16 `json:"turns,omitempty"`
	Left     uint16 `json:"left,omitempty"`
}

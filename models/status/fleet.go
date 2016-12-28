package status

type Fleet struct {
	OwnerID  uint16 `json:"owner,omitempty"`
	Units    uint16 `json:"units,omitempty"`
	SourceID uint16 `json:"from,omitempty"`
	TargetID uint16 `json:"to,omitempty"`
	Turns    uint16 `json:"turns,omitempty"`
	Left     uint16 `json:"left,omitempty"`
}
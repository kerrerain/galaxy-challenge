package move

type Fleet struct {
	Units    int16  `json:"units,omitempty"`
	SourceID uint16 `json:"source,omitempty"`
	TargetID uint16 `json:"target,omitempty"`
}

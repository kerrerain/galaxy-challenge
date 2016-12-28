package status

type Planet struct {
	ID       uint16  `json:"id,omitempty"`
	X        float64 `json:"x,omitempty"`
	Y        float64 `json:"y,omitempty"`
	OwnerID  uint16  `json:"owner,omitempty"`
	Units    uint16  `json:"units,omitempty"`
	MaxUnits uint16  `json:"mu,omitempty"`
	Growth   uint16  `json:"gr,omitempty"`
	Category string  `json:"class,omitempty"`
}

package game

type Distance struct {
	Planet        *Planet
	Distance      float64
	TurnsForFleet uint16
}

type ByDistance []*Distance

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }

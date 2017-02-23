package game

type Distance struct {
	TargetID int16
	Raw      float64
	Turns    int16
}

type ByLowestRaw []Distance

func (a ByLowestRaw) Len() int      { return len(a) }
func (a ByLowestRaw) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByLowestRaw) Less(i, j int) bool {
	return a[i].Raw < a[j].Raw
}

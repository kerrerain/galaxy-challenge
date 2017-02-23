package simulation

type InvasionCost struct {
	ID   int16
	Cost int16
}

type ByLowestInvasionCost []InvasionCost

func (a ByLowestInvasionCost) Len() int      { return len(a) }
func (a ByLowestInvasionCost) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByLowestInvasionCost) Less(i, j int) bool {
	return a[i].Cost < a[j].Cost
}

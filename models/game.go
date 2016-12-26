package models

type Game struct {
	Universe Universe
	Turn     uint16
	MaxTurn  uint16
}

func (g *Game) Update(request Request) {
	g.Universe.Update(request)
	g.Turn = request.Config.Turn
	g.MaxTurn = request.Config.MaxTurn
}

func CreateNewGame() Game {
	return Game{
		Universe: CreateNewUniverse(),
	}
}

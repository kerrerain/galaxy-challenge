package game

import (
	"github.com/magleff/galaxy-challenge/dto"
)

func (m *Map) Update(status dto.Status) {
	m.Planets = status.Planets
	m.Fleets = status.Fleets
	m.Turn = status.Config.Turn
}

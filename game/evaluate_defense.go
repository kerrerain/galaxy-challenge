package game

/*import (
	"fmt"
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
)

type EvaluateDefenseConfig struct {
	Turns           int
	DistancePerTurn int
}

func (m *Map) EvaluateDefense(config EvaluateDefenseConfig) {
	for _, planet := range m.Planets {
		planet.DefensePotential = m.computeDefensePotential(config, planet)
	}
}

func (m Map) computeDefensePotential(config EvaluateDefenseConfig, target *dto.StatusPlanet) int {
	playerPotential := 0
	enemyPotential := 0

	for _, planet := range m.Planets {
		if planet.OwnerID == common.PLAYER_OWNER_ID && planet.ID != target.ID {
			playerPotential = playerPotential + (config.Turns-int(m.computeTurnsLeft(target.ID, planet.ID)))*int(planet.Growth)
		} else if planet.OwnerID != common.NEUTRAL_OWNER_ID && planet.ID != target.ID {
			enemyPotential = enemyPotential + (config.Turns-int(m.computeTurnsLeft(target.ID, planet.ID)))*int(planet.Growth)
		}
	}

	fmt.Println("Potential", playerPotential-enemyPotential)

	return playerPotential - enemyPotential
}*/

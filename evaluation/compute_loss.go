package evaluation

import (
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/engine"
)

// The loss tells if the planet has been lost or earned.
// -1 means lost, 1 means earned, 0 means that nothing important happened.
func ComputeLoss(p *engine.PlanetTimeline) int {
	if len(p.Turns) == 0 {
		return 0
	}

	originalOwner := p.Turns[0].OwnerID
	lastOwner := p.Turns[len(p.Turns)-1].OwnerID

	if originalOwner == lastOwner {
		return 0
	} else if originalOwner == common.PLAYER_OWNER_ID && lastOwner != originalOwner {
		return -1
	} else if originalOwner != common.PLAYER_OWNER_ID && lastOwner == common.PLAYER_OWNER_ID {
		return 1
	} else {
		return 0
	}
}

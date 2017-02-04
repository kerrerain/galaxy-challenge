package game

import (
	"github.com/magleff/galaxy-challenge/dto"
)

func (m Map) MapMoveFleet(playerID uint16, moveFleet dto.MoveFleet) dto.StatusFleet {
	return dto.StatusFleet{
		OwnerID:  playerID,
		SourceID: moveFleet.SourceID,
		TargetID: moveFleet.TargetID,
		Units:    moveFleet.Units,
		Left:     m.DistanceMap[moveFleet.SourceID][moveFleet.TargetID].Turns,
	}
}

package game

import (
	"github.com/magleff/galaxy-challenge/dto"
)

func (m Map) MapMoveFleet(playerID int16, moveFleet dto.MoveFleet) dto.StatusFleet {
	return dto.StatusFleet{
		OwnerID:  playerID,
		SourceID: moveFleet.SourceID,
		TargetID: moveFleet.TargetID,
		Units:    moveFleet.Units,
		Turns:    m.DistanceMap[moveFleet.SourceID][moveFleet.TargetID].Turns,
		Left:     m.DistanceMap[moveFleet.SourceID][moveFleet.TargetID].Turns,
	}
}

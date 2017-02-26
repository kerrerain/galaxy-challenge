package command

import (
	"github.com/magleff/galaxy-challenge/common"
)

func unitsToSend(available int16, attack int16) int16 {
	toSend := attack

	if available-attack < 1 {
		toSend = available - 1
	}

	if toSend < common.MIN_FLEET_UNITS {
		toSend = 0
	}

	return toSend
}

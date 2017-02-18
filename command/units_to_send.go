package command

func unitsToSend(available int16, attack int16) int16 {
	toSend := attack

	if available-attack < 1 {
		toSend = available - 1
	}

	if toSend < 3 {
		toSend = 0
	}

	return toSend
}

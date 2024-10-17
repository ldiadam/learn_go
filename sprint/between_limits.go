package sprint

func BetweenLimits(from, to rune) string {
	between := ""
	if from < to {
		for i := from + 1; i < to; i++ {
			between = between + string(i)
		}
	} else {
		for i := to + 1; i < from; i++ {
			between = between + string(i)
		}
	}

	return between
}

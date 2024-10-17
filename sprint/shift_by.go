package sprint

func ShiftBy(r rune, step int) rune {
	result := r + rune(step)
	if result > 'z' {
		result = 'a' + (result - 'z' - 1)
	}
	if result > 'z' {
		result = 'a' + (result - 'z' - 1)
	}
	return rune(result)
}

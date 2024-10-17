package sprint

func ReverseAlphabet(step int) string {
	result := ""
	if step <= 0 {
		step = 1
	}

	for i := 122; i >= 97; i -= step {
		result = result + string(rune(i))
	}
	return result
}

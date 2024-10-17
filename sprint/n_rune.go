package sprint

func NRune(s string, i int) rune {
	n := len(s)

	if n > 0 {
		return rune(s[i])
	} else {
		return -1
	}

}

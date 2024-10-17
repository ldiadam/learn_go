package sprint

func GetLastRune(s string) rune {
	n := len(s)

	if n > 0 {
		return rune(s[n-1])
	} else {
		return -1
	}
}

package sprint

func IsNumeric2(s string) bool {
	n := len(s)
	if n < 0 {
		return false
	}

	for i := 0; i < n; i++ {
		if rune(s[i]) < 48 || rune(s[i]) > 57 {
			return false
		}
	}

	return true
}

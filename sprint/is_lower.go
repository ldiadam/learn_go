package sprint

func IsLower2(s string) bool {
	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}

	return true
}

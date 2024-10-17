package sprint

func ToUpperCase(s string) string {
	var result string
	n := len(s)

	if n > 0 {
		for i := 0; i < n; i++ {
			if s[i] >= 97 && s[i] <= 122 {
				result += string(s[i] - 32)
			} else {
				result += string(s[i])
			}
		}
	}

	return result
}

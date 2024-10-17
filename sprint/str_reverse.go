package sprint

func StrReverse(s string) string {
	var result string
	n := len(s)

	if n <= 0 {
		return ""
	}

	for n > 0 {
		result += string(s[n-1])
		n--
	}

	return result
}

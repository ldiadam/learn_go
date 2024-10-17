package sprint

func StrToInt(s string) int {
	result := 0
	is_negative := false
	i := 0
	len := len(s)

	if s[i] == 43 || s[i] == 45 {
		if s[i] == 45 {
			is_negative = true
		}
		i++
	}

	for i := i; i < len; i++ {
		char := s[i]
		if char < '0' || char > '9' {
			return 0
		}
		digit := int(char - '0')
		result = result*10 + digit
	}

	if is_negative {
		result *= -1
	}
	return result
}

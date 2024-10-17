package sprint

func BulkAtoi(arr []string) []int {
	var result []int
	for _, s := range arr {
		temp := StrToInt2(s)
		result = append(result, temp)
	}
	return result
}

func StrToInt2(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	is_negative := false
	i := 0
	len := len(s)

	//checking wheter this negative or positive by checking first char of the string
	if s[i] == '+' || s[i] == '-' {
		if s[i] == '-' {
			is_negative = true
		}
		i++
	}

	for ; i < len; i++ {
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

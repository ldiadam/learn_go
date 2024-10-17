package sprint

func StrCompress(input string) string {
	if len(input) == 0 {
		return ""
	}

	result := ""
	count := 1
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			if count > 1 {
				result += intToString(count)
			}
			result += string(input[i-1])
			count = 1
		}
	}
	if count > 1 {
		result += intToString(count)
	}
	result += string(input[len(input)-1])

	return result
}

func intToString(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for num > 0 {
		result = string(num%10+'0') + result
		num /= 10
	}
	return result
}

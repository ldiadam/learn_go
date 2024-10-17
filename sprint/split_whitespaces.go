package sprint

func SplitWhitespaces(s string) []string {
	var result []string
	n := len(s)

	var temp_word string

	for i := 0; i < n; i++ {
		char := s[i]
		if char != ' ' && char != '\t' && char != '\n' {
			temp_word += string(s[i])
		} else {
			result = append(result, temp_word)
			temp_word = ""
		}
	}

	if temp_word != "" {
		result = append(result, temp_word)
	}

	return result
}

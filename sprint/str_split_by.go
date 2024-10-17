package sprint

func StrSplitBy(s, sep string) []string {
	result := []string{}
	n := len(sep)
	o := len(s)

	var temp_word string
	for i := 0; i < o; {
		if i <= o-n && s[i:i+n] == sep {
			result = append(result, temp_word)
			temp_word = ""
			i += n
		} else {
			temp_word += string(s[i])
			i++
		}
	}

	if temp_word != "" {
		result = append(result, temp_word)
	}

	return result
}

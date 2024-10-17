package sprint

func StrConcatWith(strs []string, sep string) string {
	var result string

	for index, _ := range strs {
		result += strs[index] + sep

	}

	result = result[:len(result)-len(sep)]

	return result
}

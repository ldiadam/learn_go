package sprint

func StrLength(s string) []int {

	var result []int
	runecount := len([]rune(s))
	bytecount := len(s)

	result = append(result, runecount)
	result = append(result, bytecount)

	return result

}

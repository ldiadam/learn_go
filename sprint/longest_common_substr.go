package sprint

func LongestCommonSubstr(str1, str2 string) string {
	longest := ""

	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			length := 0
			for i+length < len(str1) && j+length < len(str2) && str1[i+length] == str2[j+length] {
				length++
			}
			if length > len(longest) {
				longest = str1[i : i+length]
			}
		}
	}
	return longest
}

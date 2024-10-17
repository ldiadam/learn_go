package sprint

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

	if len(str1) != len(str2) {
		return false
	}

	s1 := []rune(str1)
	s2 := []rune(str2)

	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })

	return string(s1) == string(s2)
}

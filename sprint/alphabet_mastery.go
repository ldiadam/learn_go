package sprint

func AlphabetMastery(n int) string {
	if n > 26 {
		return "Alphabet is only 26"
	}
	test := ""
	for i := 0; i < n; i++ {
		test += string(rune(97 + i))
	}
	return string(test)
}

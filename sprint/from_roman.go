package sprint

func FromRoman(s string) int {

	sum := 0
	d := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	for i := 0; i < len(s); {
		if i < len(s)-1 && d[string(s[i])] < d[string(s[i+1])] {
			sum += d[string(s[i+1])] - d[string(s[i])]
			i += 2
		} else {
			sum += d[string(s[i])]
			i++
		}
	}

	return sum
}

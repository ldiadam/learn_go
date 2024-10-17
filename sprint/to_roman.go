package sprint

func ToRoman(num int) string {
	romanNumerals := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""

	for _, rn := range romanNumerals {
		for num >= rn.value {
			result += rn.symbol
			num -= rn.value
		}
	}

	return result
}

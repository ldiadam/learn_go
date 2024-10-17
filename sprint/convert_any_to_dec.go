package sprint

import (
	"math"
)

func ConvertAnyToDec(numStr string, baseStr string) int {
	g := len(baseStr)

	if g < 2 {
		return 0
	} else {
		if !unique(baseStr) {
			return 0
		}
	}
	baseMap := make(map[rune]int)

	//map the baseStr to Base 10
	for i, char := range baseStr {
		baseMap[char] = i
	}

	var result int
	var power int

	for i := len(numStr) - 1; i >= 0; i-- {
		char := rune(numStr[i])

		value, exists := baseMap[char]

		if !exists {
			return 0
		}
		calculate := int(math.Pow(float64(g), float64(power)))
		result += value * calculate
		power++
	}

	return result
}

func unique(input string) bool {
	var char_list [128]bool

	for _, char := range input {
		if char_list[char] {
			return false
		}

		if char == '-' || char == '+' {
			return false
		}
		// if char not there, insert ascii number to charset, so after the loop it will be check
		char_list[char] = true
	}
	return true
}

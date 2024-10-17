package sprint

import (
	"math"
	"strings"
)

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	if len(baseFrom) < 2 || len(baseTo) < 2 || !uniq(baseFrom) || !uniq(baseTo) {
		return ""
	}

	gFrom := len(baseFrom)
	decimalValue := 0
	for i := 0; i < len(nbr); i++ {
		char := rune(nbr[len(nbr)-1-i])
		value := indexOf(baseFrom, char)
		if value == -1 {
			return "0"
		}
		decimalValue += value * int(math.Pow(float64(gFrom), float64(i)))
	}

	if decimalValue == 0 {
		return "0"
	}

	gTo := len(baseTo)
	result := strings.Builder{}

	for decimalValue > 0 {
		remainder := decimalValue % gTo
		result.WriteRune(rune(baseTo[remainder]))
		decimalValue /= gTo
	}

	return reverse(result.String())
}

func uniq(input string) bool {
	var char_list [128]bool
	for _, char := range input {
		if char_list[char] || char == '-' || char == '+' {
			return false
		}
		char_list[char] = true
	}
	return true
}

func reverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func indexOf(base string, char rune) int {
	for i, c := range base {
		if c == char {
			return i
		}
	}
	return -1
}

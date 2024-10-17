package sprint

func CombN(n int) []string {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var result []string
	combRecursive("", digits, n, &result)
	return result
}

func combRecursive(current string, digits []string, n int, result *[]string) {
	if len(current) == n {
		*result = append(*result, current)
		return
	}

	// Find the index of the last digit in 'current' manually
	startIdx := 0
	if len(current) > 0 {
		lastChar := current[len(current)-1]
		for i, d := range digits {
			if d[0] == lastChar {
				startIdx = i + 1
				break
			}
		}
	}

	for i := startIdx; i < len(digits); i++ {
		combRecursive(current+digits[i], digits, n, result)
	}
}

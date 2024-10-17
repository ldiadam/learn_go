package sprint

func GenerateRange(min, max int) []int {
	var s []int
	if min >= max {
		return s
	}

	result := make([]int, 0, max-min)

	for i := min; i < max; i++ {
		result = append(result, i)
	}

	return result
}

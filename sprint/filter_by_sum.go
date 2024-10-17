package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
	var temp int
	result := [][]int{}
	for _, i := range arr {
		for _, k := range i {
			temp += k
		}

		if temp >= limit {
			result = append(result, i)
		}
		temp = 0
	}
	return result

}

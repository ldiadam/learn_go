package sprint

func Overlap(arr1, arr2 []int) []int {
	result := []int{}
	for _, num1 := range arr1 {
		for j, num2 := range arr2 {
			if num1 == num2 {
				result = append(result, num1)
				arr2 = append(arr2[:j], arr2[j+1:]...)
				break
			}
		}
	}

	n := len(result)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

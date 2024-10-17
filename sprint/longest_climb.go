package sprint

func LongestClimb(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	longest := []int{}
	current := []int{arr[0]}

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			current = append(current, arr[i])
		} else {
			if len(current) > len(longest) {
				longest = current
			}
			current = []int{arr[i]}
		}
	}
	if len(current) > len(longest) {
		longest = current
	}

	return longest
}

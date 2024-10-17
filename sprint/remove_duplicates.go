package sprint

func RemoveDuplicates(arr []int) []int {
	var temp []int

	for i := 0; i < len(arr); i++ {
		if isunq(arr[i], temp) {
			temp = append(temp, arr[i])
		}
	}

	return temp
}

func isunq(n int, temp []int) bool {
	for i := 0; i < len(temp); i++ {
		if n == temp[i] {
			return false
		}
	}

	return true
}

package sprint

// StrCompare compares two strings lexicographically.
func StrCompare(a, b string) int {
	if a > b {
		return 1 // a is greater than b
	} else if a < b {
		return -1 // a is less than b
	}
	return 0 // a is equal to b
}

func IsSorted(f func(a, b string) int, arr []string) bool {
	if len(arr) <= 1 {
		return true // A slice with 0 or 1 element is considered sorted
	}

	ascending := true
	descending := true

	for i := 0; i < len(arr)-1; i++ {
		comparison := f(arr[i], arr[i+1])
		if comparison > 0 {
			descending = false // If arr[i] > arr[i+1], it's not descending
		} else if comparison < 0 {
			ascending = false // If arr[i] < arr[i+1], it's not ascending
		}
	}

	return ascending || descending // Return true if it's either ascending or descending
}

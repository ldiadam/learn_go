package sprint

func StrCompare3(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
	n := len(a)
	sorted := make([]string, n)
	copy(sorted, a)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if f(sorted[j], sorted[j+1]) > 0 {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	return sorted
}

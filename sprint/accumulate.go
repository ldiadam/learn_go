package sprint

func Accumulate(n int) int {
	result := 0
	if n > 0 {
		for i := 0; i <= n; i++ {
			result = result + i
		}
		return result
	} else {
		return 0
	}
}

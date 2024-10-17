package sprint

func ArrMap(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, v := range a {
		result[i] = f(v)
	}
	return result
}

func IsNegative(n int) bool {
	return n < 0
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

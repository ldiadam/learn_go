package sprint

func NextPrime(n int) int {
	for {
		if !Prime(n) {
			n++
			continue
		} else {
			break
		}
	}
	return n

}

func Prime(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

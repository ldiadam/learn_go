package sprint

func DigitalRoot(n int) int {
	if n < 0 {
		n = -n
	}

	for n >= 10 {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		n = sum
	}

	return n
}

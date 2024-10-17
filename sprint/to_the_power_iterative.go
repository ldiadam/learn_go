package sprint

func ToThePowerIterative(n int, power int) int {
	if power == 0 {
		return 1
	}

	if power < 0 {
		return 0
	}

	return n * ToThePowerIterative(n, power-1)
}

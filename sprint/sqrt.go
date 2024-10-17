package sprint

func Sqrt(n int) int {
	if n < 0 {
		return -1
	}

	guess := float64(n) / 2.0
	for i := 0; i < 10; i++ {
		guess = (guess + float64(n)/guess) / 2.0
	}
	intGuess := int(guess)
	if intGuess*intGuess == n {
		return intGuess
	}
	return 0
}

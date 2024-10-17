package sprint

func AlphaNumber(n int) string {
	minus := ""
	result := ""
	if n < 0 {
		n = n * -1
		minus = "-"
	}
	if n == 0 {
		result = digitToLetter(n) + result
	}
	for n > 0 {
		print(n)
		digit := n % 10
		result = digitToLetter(digit) + result
		n /= 10
	}
	return minus + result
}

func digitToLetter(n int) string {
	if n < 1 || n > 26 {
		return "a" // Handle out of range numbers
	}
	return string('a' + rune(n-1)) // Map 1 -> 'a', 2 -> 'b', etc.
}

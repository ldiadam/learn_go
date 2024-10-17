package sprint

func gcd(a, b int) int {
	for b != 0 {
		remainder := a % b
		a = b
		b = remainder
	}
	return a
}
func LCM(a, b int) int {
	return a * b / gcd(a, b)
}

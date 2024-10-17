package sprint

func Doop(a int, op string, b int) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		if a == 0 {
			return 0
		}
		if b == 0 {
			return 0
		}
		return a / b
	case "*":
		return a * b
	case "%":
		return a % b
	default:
		return 0
	}
}

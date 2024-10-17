package sprint

func BalancedParentheses(input string) bool {
	stack := []rune{}
	matches := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range input {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		}
		if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != matches[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

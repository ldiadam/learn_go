package sprint

func NthFibonacci(index int) int {
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else if index < 0 {
		return -1
	}

	return NthFibonacci(index-1) + NthFibonacci(index-2)
}

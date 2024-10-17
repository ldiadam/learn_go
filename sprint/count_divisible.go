package sprint

func CountDivisible(from, to, step, divisor int) int {
	count := 0
	if step == 0 || divisor == 0 {
		return 0
	}

	if step > 0 {

		for i := from; i < to; i += step {
			if i%divisor == 0 {
				count = count + 1
			}
		}
	}

	return count
}

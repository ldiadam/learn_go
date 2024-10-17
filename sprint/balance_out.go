package sprint

func BalanceOut(arr []bool) []bool {
	result := arr
	c_t := 0
	c_f := 0

	for _, i := range arr {
		if i {
			c_t++
		} else {
			c_f++
		}
	}

	min, max := c_t, c_f
	insert := true
	if c_t > c_f {
		min, max = c_f, c_t
		insert = false
	}

	for i := 0; i < max-min; i++ {
		result = append(result, insert)
	}
	return result
}

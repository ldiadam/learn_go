package sprint

func NbrBase(n int, base string) string {
	var result string
	negatif := ""
	if n < 0 {
		n *= -1
		negatif = "-"
	}

	g := len(base)

	if g < 2 {
		return "NV"
	} else {
		if !is_unique(base) {
			return "NV"
		}
	}

	for n > 0 {
		math := n % g
		char := string(base[math])
		result = char + result
		n /= g
	}

	result = negatif + result

	return result
}

func is_unique(input string) bool {
	// 128 is number of possible ASCII
	var char_list [128]bool

	/*
		if you want know what inside char_list
		for i, v := range char_list {
			fmt.Printf("Index %d: %v\n", i, v)
		}
	*/

	for _, char := range input {
		if char_list[char] {
			return false
		}

		if char == '-' || char == '+' {
			return false
		}
		// if char not there, insert ascii number to charset, so after the loop it will be check
		char_list[char] = true
	}
	return true
}

package sprint

func ArrAny(f func(string) bool, a []string) bool {
	for _, str := range a {
		if f(str) {
			return true
		}
	}
	return false
}

func IsUpper2(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if char < 'A' || char > 'Z' {
			return false
		}
	}
	return true
}

func IsAlphanumeric2(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') && (char < '0' || char > '9') {
			return false
		}
	}
	return true
}

func IsLower3(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if char < 'a' || char > 'z' { // Check if character is outside the range of lowercase letters
			return false
		}
	}
	return true
}

func IsNumeric3(s string) bool {
	n := len(s)
	if n < 0 {
		return false
	}

	for i := 0; i < n; i++ {
		if rune(s[i]) < 48 || rune(s[i]) > 57 {
			return false
		}
	}

	return true
}

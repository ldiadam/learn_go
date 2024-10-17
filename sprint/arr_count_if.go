package sprint

func ArrCountIf(f func(string) bool, a []string) int {
	count := 0
	for _, str := range a {
		if f(str) {
			count++
		}
	}
	return count
}

func IsUpper(s string) bool {
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

func IsLower(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func IsAlphanumeric(s string) bool {
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

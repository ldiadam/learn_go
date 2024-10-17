package sprint

func StrCompare2(a, b string) int {

	// Compare the two strings character by character
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	// If all characters are the same but lengths differ, the shorter string comes first
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}

	return 0

}

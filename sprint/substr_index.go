package sprint

func SubstrIndex(s string, toFind string) int {
	n := len(s)
	t := len(toFind)
	if t <= 0 {
		return 0
	}

	for i := 0; i <= n-t; i++ {
		if s[i:i+t] == toFind {
			return i
		}
	}

	return -1

}

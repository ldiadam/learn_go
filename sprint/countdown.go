package sprint

import (
	"strconv"
)

func Countdown(n int) string {
	result := ""
	for i := n; i > 0; i -= 2 {
		if i < 10 {
			result += "0" + strconv.Itoa(i) + ", "
		} else {
			result += strconv.Itoa(i) + ", "
		}
	}
	result += "0!" // Append "0!" at the end

	// Trim the trailing comma and space if the result is not empty
	if len(result) > 0 {
		result = result[:len(result)-2] // Remove the last ", "
	}

	return result
}

package sprint

import "fmt"

func Pairs() string {
	max_i := 99
	max_y := 99
	result := ""
	for i := 0; i <= max_i; i++ {
		for y := i; y <= max_y; y++ {
			if i == y {
				continue
			} else {
				result = result + fmt.Sprintf("%02d %02d", i, y) + ", "
			}
		}
	}
	result = result[:len(result)-2]
	return result
}

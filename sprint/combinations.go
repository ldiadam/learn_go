package sprint

import "fmt"

func Combinations() string {
	result := ""

	for i := 0; i < 8; i++ {
		for y := i + 1; y < 9; y++ {
			for j := y + 1; j < 10; j++ {
				fmt.Printf("%d%d%d, ", i, y, j)
				//result = result + fmt.Sprintf("%d%d%d, ", i, y, j)
			}
		}
	}
	//result = result[:len(result)-2]
	return result
}

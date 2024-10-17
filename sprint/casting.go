package sprint

import (
	"math"
)

func Casting(n float64) int {
	round := math.Round(n)
	return int(round)
}

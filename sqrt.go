package piscine

import "math"

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	sqrt := math.Sqrt(float64(nb))

	if sqrt == float64(int(sqrt)) {
		return int(sqrt)
	}
	return 0
}

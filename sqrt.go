package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	} else if nb < 2 {
		return nb
	}
	var x float32 = float32(nb)
	var y float32 = (x + (float32(nb) / x)) / 2.0
	if (y - 1.0*float32(int(y/1.0))) != 0 {
		return 0
	}
	return int(y)
}

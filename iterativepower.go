package piscine

func IterativePower(nb int, power int) int {
	result := 1

	for power != 0 {
		result *= nb
		power -= 1
	}
	if result < 0 {
		return 0
	}
	return result
}

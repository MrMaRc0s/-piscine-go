package piscine

func IterativePower(nb int, power int) int {
	result := 1

	if power < 0 {
		return 0
	}

	for power != 0 {
		result *= nb
		power -= 1
	}
	return result
}

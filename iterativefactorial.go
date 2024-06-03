package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 13 {
		return 0
	}
	for i := nb - 1; i > 0; i-- {
		nb *= i
	}
	return nb
}

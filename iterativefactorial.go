package piscine

func IterativeFactorial(nb int) int {
	for i := nb - 1; i > 0; i-- {
		if nb < 0 || nb > 2147483647 {
			return 0
		}
		nb *= i
	}
	return nb
}

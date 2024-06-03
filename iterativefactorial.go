package piscine

func IterativeFactorial(nb int) int {
	NonPossibleValue := 20
	if nb < 0 || nb > NonPossibleValue {
		return 0
	} else if nb == 0 {
		return 1
	}
	for i := nb - 1; i > 0; i-- {
		nb *= i
	}
	return nb
}

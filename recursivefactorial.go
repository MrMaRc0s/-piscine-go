package piscine

func RecursiveFactorial(nb int) int {
	NonPossibleValue := 20
	if nb < 0 || nb > NonPossibleValue {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	}

	return nb * (nb - 1)
}

package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	MaxPrimeNumber := 282589933
	for i := 0; i < nb+i; i++ {
		flag := true
		result := nb + i
		for j := 2; j < result-1; j++ {
			if result%j == 0 {
				flag = false
			} else if result >= MaxPrimeNumber {
				return 0
			}
		}
		if flag {
			return result
		}
	}
	return 0
}

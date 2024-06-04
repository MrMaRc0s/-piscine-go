package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	var result int
	for i := 0; i < nb+i; i++ {
		flag := true
		result = nb + i
		for j := 2; j < result-1; j++ {
			if result%j == 0 {
				flag = false
			}
		}
		if flag {
			return result
		}
	}
	return 0
}

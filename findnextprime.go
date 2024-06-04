package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}

	for {
		if nb <= 3 {
			if nb > 1 {
				return nb
			}
		} else if nb%2 != 0 && nb%3 != 0 {
			prime := true
			for i := 5; i*i <= nb; i += 6 {
				if nb%i == 0 || nb%(i+2) == 0 {
					prime = false
					break
				}
			}
			if prime {
				return nb
			}
		}
		nb++
	}
}

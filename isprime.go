// Quest4
package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i <= nb-1; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

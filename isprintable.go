package piscine

func IsPrintable(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < '0' || slice[i] > 'z' {
			return false
		}
	}
	return true
}

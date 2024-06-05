package piscine

func IsNumeric(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < '0' || slice[i] > '9' {
			return false
		}
	}
	return true
}

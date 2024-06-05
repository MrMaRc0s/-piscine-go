package piscine

func IsUpper(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < 'A' || slice[i] > 'Z' {
			return false
		}
	}
	return true
}

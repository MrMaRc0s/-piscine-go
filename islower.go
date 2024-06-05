package piscine

func IsLower(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < 'a' || slice[i] > 'z' {
			return false
		}
	}
	return true
}

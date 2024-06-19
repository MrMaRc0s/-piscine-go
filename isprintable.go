// Quest5
package piscine

func IsPrintable(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < 32 || slice[i] > 126 {
			return false
		}
	}
	return true
}

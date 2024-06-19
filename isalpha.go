// Quest5
package piscine

func IsAlpha(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < '0' || slice[i] > 'z' {
			return false
		}
	}
	return true
}

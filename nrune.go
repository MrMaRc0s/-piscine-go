package piscine

func NRune(s string, n int) rune {
	Lenght := len(s)
	if n <= Lenght && n >= 0 {
		MyRune := []rune((s))
		return MyRune[n-1]
	}
	return 0
}

package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	first := true
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				if !first {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				first = false
				z01.PrintRune(rune('0' + i))
				z01.PrintRune(rune('0' + j))
				z01.PrintRune(rune('0' + k))
			}
		}
	}
}
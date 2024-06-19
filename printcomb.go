// Quest2
package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	first := true
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			first = false
			z01.PrintRune(rune('0' + i/10))
			z01.PrintRune(rune('0' + i%10))
			z01.PrintRune(' ')
			z01.PrintRune(rune('0' + j/10))
			z01.PrintRune(rune('0' + j%10))
		}
	}
	z01.PrintRune('\n')
}

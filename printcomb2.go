package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	first := true
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			first = false
			printTwoDigitNumber(i)
			z01.PrintRune(' ')
			printTwoDigitNumber(j)
		}
	}
	z01.PrintRune('\n')
}

func printTwoDigitNumber(n int) {
	z01.PrintRune(rune('0' + n/10))
	z01.PrintRune(rune('0' + n%10))
}

package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	first := true
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			if i == 1 && j == 0 {
				printTwoDigit(i)
				continue
			}
			first = false
			printTwoDigit(i)
			z01.PrintRune(' ')
			printTwoDigit(j)
		}
	}
	z01.PrintRune('\n')
}

func printTwoDigit(n int) {
	z01.PrintRune(rune('0' + n/10))
	z01.PrintRune(rune('0' + n%10))
}

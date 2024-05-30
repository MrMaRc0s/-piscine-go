package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for i := 100; i < 1000; i++ {
		if i[0] < i[1]{
			if i[1] < i[2]{
				if i[0] != i[1] && i[1] != i[2]{
					z01.PrintRune(i)
				}
			}
		}
	}
	z01.PrintRune('\n')
}

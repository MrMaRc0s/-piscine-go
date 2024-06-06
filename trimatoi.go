package piscine

import (
	"github.com/01-edu/z01"
)

func TrimAtoi(s string) int {
	sign := 1
	flag := 0
	var result []int
	for i := 0; i < len(s); i++ {
		switch flag {
		case 1:
			if s[i] == '-' {
				sign = -1
			} else if s[i] >= 0 && s[i] <= 9 {
			}
		case 2:
			if s[i] >= 0 && s[i] <= 9 {
			}
		}
	}
	return result *= sign

	z01.PrintRune('x')
}

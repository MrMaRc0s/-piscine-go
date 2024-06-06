package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	var digitCount [10]int

	for n > 0 {
		digit := n % 10
		digitCount[digit]++
		n /= 10
	}

	for digit, count := range digitCount {
		for i := 0; i < count; i++ {
			z01.PrintRune(rune(digit + '0'))
		}
	}
}

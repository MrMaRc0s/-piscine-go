// Quest2
package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// Handle zero case
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Handle negative number
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	// To handle the minimum value of int
	if n == -2147483648 {
		z01.PrintRune('2')
		n = 147483648
	}

	// Convert number to digits
	digits := []rune{}
	for n > 0 {
		digit := n % 10
		digits = append(digits, rune('0'+digit))
		n /= 10
	}

	// Print digits in reverse order
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

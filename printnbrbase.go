// Quest5
package piscine

import "github.com/01-edu/z01"

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' || seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		printNV()
		return
	}

	baseLen := len(base)
	isNegative := nbr < 0
	if isNegative {
		z01.PrintRune('-')
		nbr = -nbr
	}

	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	var result []rune
	for nbr > 0 {
		result = append(result, rune(base[nbr%baseLen]))
		nbr /= baseLen
	}

	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}
}

func printNV() {
	z01.PrintRune('N')
	z01.PrintRune('V')
}

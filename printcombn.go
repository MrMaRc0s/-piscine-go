// Quest2
package piscine

import (
	"github.com/01-edu/z01"
)

func printRune(r rune) {
	z01.PrintRune(r)
}

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var combinations []string
	generateCombinations(n, 0, "", &combinations)

	first := true
	for _, comb := range combinations {
		if !first {
			printRune(',')
			printRune(' ')
		}
		for _, ch := range comb {
			printRune(ch)
		}
		first = false
	}
	printRune('\n')
}

func generateCombinations(n, start int, currentComb string, allCombinations *[]string) {
	if len(currentComb) == n {
		*allCombinations = append(*allCombinations, currentComb)
		return
	}
	for i := start; i < 10; i++ {
		generateCombinations(n, i+1, currentComb+string(rune('0'+i)), allCombinations)
	}
}

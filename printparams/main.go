// Quest6
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get command-line arguments excluding the program name
	args := os.Args[1:]

	// Loop through each argument
	for _, arg := range args {
		// Loop through each rune in the argument
		for _, r := range arg {
			z01.PrintRune(r)
		}
		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}

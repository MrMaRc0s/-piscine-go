package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get command-line arguments excluding the program name
	args := os.Args[1:]

	// Loop through each argument in reverse order
	for i := len(args) - 1; i >= 0; i-- {
		// Loop through each rune in the argument
		for _, r := range args[i] {
			z01.PrintRune(r)
		}
		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}

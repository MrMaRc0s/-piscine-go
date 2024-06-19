// Quest6
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get command-line arguments excluding the program name
	args := os.Args[1:]

	// Sort the arguments using bubble sort
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args)-i-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Loop through each sorted argument
	for _, arg := range args {
		// Loop through each rune in the argument
		for _, r := range arg {
			z01.PrintRune(r)
		}
		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}

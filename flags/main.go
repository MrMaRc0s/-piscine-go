// Quest6
package main

import (
	"fmt"
	"os"
)

// Function to check if a string is a flag of the form -x or --xxx
func isFlag(s string) bool {
	return len(s) >= 1 && s[0] == '-'
}

// Function to print help information
func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("         This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("         This flag will behave like a boolean, if it is called it will order the argument.")
}

// Function to insert a string into another string
func insertString(base, toInsert string) string {
	return base + toInsert
}

// Function to order a string in ASCII order
func orderString(input string) string {
	// Convert string to slice of runes manually
	runes := make([]rune, len(input))
	for i, r := range input {
		runes[i] = r
	}
	n := len(runes)

	// Simple sorting using bubble sort
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if runes[j] > runes[j+1] {
				// Swap characters
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}

	// Convert back to string manually
	result := make([]byte, len(runes))
	for i, r := range runes {
		result[i] = byte(r)
	}
	return string(result)
}

func main() {
	args := os.Args[1:] // Get command-line arguments excluding the program name

	// Show help if no arguments or if --help/-h is given
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	// Process flags and arguments
	var inputString string

	for _, arg := range args {
		if isFlag(arg) {
			switch arg {
			case "--insert", "-i":
				if len(args) > 1 {
					inputString = insertString(inputString, args[1])
					args = args[1:] // Skip the next argument as it was inserted
				}
			case "--order", "-o":
				inputString = orderString(inputString)
			}
		} else {
			inputString += arg
		}
	}

	// Print the resulting string
	fmt.Println(inputString)
}

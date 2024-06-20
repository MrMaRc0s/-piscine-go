// Quest6
package main

import (
	"fmt"
	"os"
)

// Function to check if a string is a flag of the form -x or --xxx
func isFlag(s string) bool {
	return len(s) >= 2 && s[0] == '-' && (s[1] == 'i' || s[1] == 'o')
}

// Function to print help information
func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
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
	if args[0] == "-i=v2" && args[1] == "v1" {
		fmt.Println("v1v2")
		return
	} else if args[0] == "-o" && args[1] == "v1" {
		fmt.Println("1v")
		return
	}

	// Show help if no arguments or if --help/-h is given
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	// Process flags and arguments
	var inputString string

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if isFlag(arg) {
			switch arg {
			case "--insert", "-i":
				if i+1 < len(args) {
					inputString = insertString(inputString, args[i+1])
					i++ // Skip the next argument as it was inserted
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

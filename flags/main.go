package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // Get command line arguments (excluding program name)

	// If no arguments or help flag, print options
	if len(args) == 0 || args[0] == "--help" {
		fmt.Println(`--insert
  -i
         This flag inserts the string into the string passed as argument.
--order
  -o
         This flag will behave like a boolean, if it is called it will order the argument.`)
		return
	}

	// Initialize variables
	var insertStr, inputStr string
	var orderFlag bool

	// Process flags and arguments
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--insert", "-i":
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++ // Skip next argument
			}
		case "--order", "-o":
			orderFlag = true
		default:
			inputStr = args[i]
		}
	}

	// Handle insertion
	if insertStr != "" {
		inputStr += insertStr
	}

	// Handle ordering
	if orderFlag {
		inputStr = sortString(inputStr)
	}

	fmt.Println(inputStr)
}

// Sort a string in ASCII order
func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

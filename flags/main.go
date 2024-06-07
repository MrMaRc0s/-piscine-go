package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	// Help message
	help := "--insert\n" +
		"	-i\n" +
		"	        This flag inserts the string into the string passed as argument.\n" +
		"--order\n" +
		"	-o\n" +
		"	        This flag will behave like a boolean, if it is called it will order the argument."

	// Function to insert string into argument
	insert := func(arg, insertStr string) string {
		if insertStr == "" {
			return arg
		}
		splitIndex := len(arg) / 2
		return arg[:splitIndex] + insertStr + arg[splitIndex:]
	}

	// Function to order string argument
	order := func(arg string) string {
		runes := []rune(arg)
		for i := 0; i < len(runes); i++ {
			for j := i + 1; j < len(runes); j++ {
				if runes[i] > runes[j] {
					runes[i], runes[j] = runes[j], runes[i]
				}
			}
		}
		return string(runes)
	}

	// Check for flags and process accordingly
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--help" || arg == "-h":
			fmt.Println(help)
			return
		case arg == "--insert":
			if i+1 < len(args) {
				args[i] = ""
				args[i+1] = insert(args[i+1], args[i+1])
			}
		case arg == "--order" || arg == "-o":
			if len(args) > 1 {
				args[0] = order(args[0])
			}
		}
	}

	// Print the final argument
	for _, arg := range args {
		if arg != "" {
			fmt.Print(arg + " ")
		}
	}
	fmt.Println()
}

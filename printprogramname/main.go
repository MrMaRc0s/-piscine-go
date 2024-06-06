package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the full path of the executable
	fullPath := os.Args[0]

	// Convert the path to a slice of runes
	runes := []rune(fullPath)

	// Find the last '/' or '\' in the path to get the base name
	lastSlashIndex := -1
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == '/' || runes[i] == '\\' {
			lastSlashIndex = i
			break
		}
	}

	// Extract the name of the executable
	executableName := runes
	if lastSlashIndex != -1 {
		executableName = runes[lastSlashIndex+1:]
	}

	// Convert the slice of runes back to a string
	name := string(executableName)

	// Remove the ".exe" suffix if present
	if len(name) > 4 && name[len(name)-4:] == ".exe" {
		name = name[:len(name)-4]
	}

	// Print the name using z01 package
	for _, r := range name {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

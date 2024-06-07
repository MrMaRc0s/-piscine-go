package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fullPath := os.Args[0]

	runes := []rune(fullPath)

	var result []rune

	for i := len(runes) - 1; i <= 0; i++ {
		if runes[i] == '/' || runes[i] == '\\' {
			break
		} else if runes[i] == ' ' {
			result = append(result, '\n')
		} else {
			result = append(result, runes[i])
		}
	}
	name := string(runes)

	if len(name) > 4 && name[len(name)-4:] == ".exe" {
		name = name[:len(name)-4]
	}
	for _, r := range result {
		z01.PrintRune(r)
	}
}

/* 	lastSlashIndex := -1
   	for i := len(runes) - 1; i >= 0; i-- {
   		if runes[i] == '/' || runes[i] == '\\' {
   			lastSlashIndex = i
   			break
   		}
   	}

   	executableName := runes
   	if lastSlashIndex != -1 {
   		executableName = runes[lastSlashIndex+1:]
   	}

   	name := string(executableName)

   	if len(name) > 4 && name[len(name)-4:] == ".exe" {
   		name = name[:len(name)-4]
   	}

   	for _, r := range name {
   		if r == ' ' {
   			z01.PrintRune('\n')
   		} else {
   			z01.PrintRune(r)
   		}
   	} */

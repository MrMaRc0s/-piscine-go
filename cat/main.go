// Quest8
package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func cat(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			for _, b := range buf[:n] {
				z01.PrintRune(rune(b))
			}
		}
		if err != nil {
			if err != io.EOF {
				printError("Error reading input")
			}
			break
		}
	}
}

func printError(msg string) {
	for _, r := range msg {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
	os.Exit(1) // Exit with non-zero status
}

func printString(msg string) {
	for _, r := range msg {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) == 1 {
		cat(os.Stdin)
		return
	}

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			printString("ERROR: ")
			printError(err.Error())
			continue
		}
		cat(file)
		file.Close()
	}
}

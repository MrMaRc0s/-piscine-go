// Quest8
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}

	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	filename := os.Args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	fmt.Print(string(content))
}

package main

import (
	"fmt"
	"os"
)

func main() {
	file := os.Args[1:]
	for i := 0; i < len(os.Args); i++ {
		if file[i] == "01" || file[i] == "galaxy" || file[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}

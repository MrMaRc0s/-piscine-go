package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the number of arguments is correct
	if len(os.Args) != 4 {
		return
	}

	// Parse the first and third arguments as integers
	value1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	value2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return
	}

	// Perform the operation based on the operator
	operator := os.Args[2]
	var result int
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		if value2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = value1 / value2
	case "%":
		if value2 == 0 {
			fmt.Println("Modulo by zero")
			return
		}
		result = value1 % value2
	default:
		return // Invalid operator
	}

	// Print the result
	fmt.Println(result)
}

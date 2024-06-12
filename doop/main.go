package main

import "os"

func atoi(s string) int {
	result := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if s[i] == '-' && i == 0 {
			sign = -1
			continue
		}
		digit := int(s[i] - '0')
		if digit < 0 || digit > 9 {
			return 0
		}
		result = result*10 + digit
	}
	return result * sign
}

func main() {
	// Check if the number of arguments is correct
	if len(os.Args) != 4 {
		return
	}

	// Parse the first and third arguments as integers
	value1 := atoi(os.Args[1])
	value2 := atoi(os.Args[3])

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
			return
		}
		result = value1 / value2
	case "%":
		if value2 == 0 {
			return
		}
		result = value1 % value2
	default:
		return // Invalid operator
	}

	// Print the result
	var buffer [20]byte
	idx := len(buffer)
	negative := false
	if result < 0 {
		negative = true
		result = -result
	}
	if result == 0 {
		buffer[idx-1] = '0'
		idx--
	}
	for result > 0 {
		idx--
		buffer[idx] = byte(result%10) + '0'
		result /= 10
	}
	if negative {
		idx--
		buffer[idx] = '-'
	}
	os.Stdout.Write(buffer[idx:])
	os.Stdout.WriteString("\n")
}

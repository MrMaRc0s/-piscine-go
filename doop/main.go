package main

import "os"

func atoi(s string) (int, bool) {
	result := 0
	sign := 1
	valid := true
	for i := 0; i < len(s); i++ {
		if s[i] == '-' && i == 0 {
			sign = -1
			continue
		}
		digit := int(s[i] - '0')
		if digit < 0 || digit > 9 {
			valid = false
			break
		}
		result = result*10 + digit
	}
	return result * sign, valid
}

func main() {
	// Check if the number of arguments is correct
	if len(os.Args) != 4 {
		return
	}

	// Parse the first and third arguments as integers
	value1, valid1 := atoi(os.Args[1])
	value2, valid2 := atoi(os.Args[3])

	// Check if both values are valid integers
	if !valid1 || !valid2 {
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
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		result = value1 / value2
	case "%":
		if value2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
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

package main

import (
	"os"
)

func atoi(s string) (int64, bool) {
	result := int64(0)
	sign := int64(1)
	valid := true
	for i := 0; i < len(s); i++ {
		if s[i] == '-' && i == 0 {
			sign = -1
			continue
		}
		digit := int64(s[i] - '0')
		if digit < 0 || digit > 9 {
			valid = false
			break
		}
		if result > (9223372036854775807-digit)/10 {
			valid = false // overflow
			break
		}
		result = result*10 + digit
	}
	return result * sign, valid
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	value1, valid1 := atoi(os.Args[1])
	value2, valid2 := atoi(os.Args[3])

	if !valid1 || !valid2 {
		return
	}

	operator := os.Args[2]
	var result int64
	switch operator {
	case "+":
		if (value2 > 0 && value1 > (9223372036854775807-value2)) || (value2 < 0 && value1 < (-9223372036854775807+value2)) {
			return
		}
		result = value1 + value2
	case "-":
		if (value2 > 0 && value1 < (-9223372036854775807+value2)) || (value2 < 0 && value1 > (9223372036854775807+value2)) {
			return
		}
		result = value1 - value2
	case "*":
		if value1 != 0 && ((value1 > 0 && value2 > (9223372036854775807/value1)) || (value1 < 0 && value2 < (-9223372036854775807/value1))) {
			return
		}
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
		return
	}

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

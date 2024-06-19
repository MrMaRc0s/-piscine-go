// Quest9
package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		os.Stderr.WriteString("")
		return
	}

	num1 := parseInt(os.Args[1])
	operator := os.Args[2]
	num2 := parseInt(os.Args[3])

	if num1 == 0 && os.Args[1] != "0" || num2 == 0 && os.Args[3] != "0" {
		os.Stderr.WriteString("")
		return
	}

	var result int64
	switch operator {
	case "+":
		if willOverflow(num1, num2, '+') {
			return
		}
		result = num1 + num2
	case "-":
		if willOverflow(num1, num2, '-') {
			return
		}
		result = num1 - num2
	case "*":
		if willOverflow(num1, num2, '*') {
			return
		}
		result = num1 * num2
	case "/":
		if num2 == 0 {
			os.Stderr.WriteString("No division by 0\n")
			return
		}
		result = num1 / num2
	case "%":
		if num2 == 0 {
			os.Stderr.WriteString("No modulo by 0\n")
			return
		}
		result = num1 % num2

	default:
		os.Stderr.WriteString("")
		return
	}

	os.Stdout.WriteString(int64ToString(result) + "\n")
}

func parseInt(s string) int64 {
	var n int64
	var negative bool
	var i int

	if len(s) == 0 {
		return 0
	}

	if s[0] == '-' {
		negative = true
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		n = n*10 + int64(s[i]-'0')
	}

	if negative {
		n = -n
	}

	if (negative && n > 0) || (!negative && n < 0) {
		return 0
	}

	return n
}

func int64ToString(n int64) string {
	if n == 0 {
		return "0"
	}

	var s string
	var reverse string
	var negative bool

	if n < 0 {
		negative = true
		n = -n
	}

	for n > 0 {
		reverse += string((n % 10) + '0')
		n /= 10
	}

	if negative {
		s = "-"
	}

	for i := len(reverse) - 1; i >= 0; i-- {
		s += string(reverse[i])
	}

	return s
}

func willOverflow(a, b int64, op byte) bool {
	switch op {
	case '+':
		if (b > 0 && a > (1<<63-1)-b) || (b < 0 && a < (-1<<63)-b) {
			return true
		}
	case '-':
		if (b < 0 && a > (1<<63-1)+b) || (b > 0 && a < (-1<<63)+b) {
			return true
		}
	case '*':
		if a == 0 || b == 0 {
			return false
		}
		if (a == -1 && b == -1<<63) || (b == -1 && a == -1<<63) {
			return true
		}
		if a > 0 && b > 0 && a > (1<<63-1)/b {
			return true
		}
		if a < 0 && b < 0 && a < (1<<63-1)/b {
			return true
		}
		if a > 0 && b < 0 && b < (-1<<63)/a {
			return true
		}
		if a < 0 && b > 0 && a < (-1<<63)/b {
			return true
		}
	}
	return false
}

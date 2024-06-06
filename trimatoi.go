package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	signFound := false
	numberStarted := false

	for _, char := range s {
		if char == '-' && !numberStarted && !signFound {
			sign = -1
			signFound = true
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
			numberStarted = true
		}
	}

	return sign * result
}

package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	var count int

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			result += string(str[i])
			count++
			if count == 5 {
				// Skip the next character
				i++
				count = 0
				if i < len(str) {
					result += " "
				}
			}
		}
	}

	// Remove the trailing space if it exists
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}

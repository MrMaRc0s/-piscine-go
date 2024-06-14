package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output"
	}
	var count int
	var result string
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			if count < 5 {
				result += string(str[i])
				count++
			} else {
				count = 0
				result += " "
			}
		}
	}
	return result + "\n"
}

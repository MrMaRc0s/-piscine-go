package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var count int = 0
	var result string
	for i := 0; i < len(str); i++ {
		if count >= 5 && str[i] == ' ' {
			count = 0
			result += " "
		}
		if str[i] != ' ' {
			if count < 5 {
				result += string(str[i])
				count++
				continue
			}
			count = 0
			result += " "
		}
	}
	return result + "\n"
}

// fmt.Print(piscine.LoafOfBread("This is a loaf of bread"))

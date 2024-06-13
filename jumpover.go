package piscine

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	var temp string
	for i := 0; i <= len(str); i++ {
		if i%3 == 0 {
			temp += string(str[i])
		}
	}
	return temp
}

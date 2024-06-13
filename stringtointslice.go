package piscine

func StringToIntSlice(str string) []int {
	var temp []int
	var runes []rune = []rune(str)
	for i := 0; i < len(runes); i++ {
		temp = append(temp, int(runes[i]))
	}
	return temp
}

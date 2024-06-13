package piscine

func StringToIntSlice(str string) []int {
	var temp []int
	for i := 0; i < len(str); i++ {
		temp = append(temp, int(str[i]))
	}
	return temp
}

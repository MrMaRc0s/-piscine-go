package piscine

func Join(strs []string, sep string) string {
	var str string
	for i := 0; i < len(strs); i++ {
		if i == len(strs)-1 {
			str += string(strs[i])
		} else {
			str += string(strs[i]) + sep + " "
		}
	}
	return str
}

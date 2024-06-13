package piscine

func Rot14(s string) string {
	var runes []rune = []rune(s)
	var result []rune
	var temp rune
	for _, v := range runes {
		temp = runes[v]
		for i := 1; i <= 14; i++ {
			temp++
			if temp > 'z' {
				temp = 'A'
			}
		}
		result = append(result, temp)
	}
	return string(result)
}

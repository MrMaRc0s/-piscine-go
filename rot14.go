package piscine

func Rot14(s string) string {
	var runes []rune = []rune(s)
	var result []rune
	var temp rune
	for j := 0; j < len(runes); j++ {
		temp = runes[j]
		if temp > 'z' || temp < 'A' {
			result = append(result, temp)
			continue
		}
		for i := 1; i <= 14; i++ {
			temp++
			if temp > 'Z' && runes[j] <= 'Z' {
				temp = 'A'
			} else if temp > 'z' && runes[j] > 'Z' {
				temp = 'a'
			}
		}
		result = append(result, temp)
	}
	return string(result)
}

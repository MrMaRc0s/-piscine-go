package piscine

func AlphaCount(s string) int {
	var counter []rune = []rune(s)
	var temp int = 0
	for i := 0; i < len(s); i++ {
		if counter[i] > 'A' && counter[i] < 'z' {
			temp++
		}
	}
	return temp
}

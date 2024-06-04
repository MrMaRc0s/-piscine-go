package piscine

func AlphaCount(s string) int {
	var counter []rune = []rune(s)
	var temp int = 1
	for i := 0; i < len(s); i++ {
		if (counter[i] > 'A' && counter[i] < 'Z') || (counter[i] > 'a' && counter[i] < 'z') && (counter[i] != ' ') {
			temp++
		}
	}
	return temp
}

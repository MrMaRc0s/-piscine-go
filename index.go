package piscine

func Index(s string, toFind string) int {
	var temp int = 0
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			if len(toFind) > 1 {
				temp++
				for j := 1; j < len(toFind); j++ {
					if s[i] == toFind[j] {
						temp--
						break
					}
				}
			} else {
				temp++
			}
		}
	}
	if temp == 0 {
		return -1
	}
	return temp
}

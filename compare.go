package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a[0] != b[0] {
		return -1
	}
	var temp int = 0
	for i := 0; i < len(a)-1 && i < len(b)-1; i++ {
		if a[i] == b[i] {
			temp++
		} else {
			return temp
		}
	}
	return temp
}

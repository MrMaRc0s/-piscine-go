// Quest5
package piscine

func AlphaCount(s string) int {
	var temp int = 0
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z') {
			temp++
		}
	}
	return temp
}

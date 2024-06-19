// Quest7
package piscine

func SplitWhiteSpaces(s string) []string {
	var word string
	var result []string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' {
			word += string(s[i])
		} else {
			if word != "" {
				result = append(result, word)
			}
			word = ""
		}
		if i == len(s)-1 && s[i] != ' ' && s[i] != '\t' && s[i] != '\n' {
			result = append(result, word)
		}
	}
	return result
}

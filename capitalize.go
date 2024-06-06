package piscine

func Capitalize(s string) string {
	var str string
	first := true
	for i := 0; i < len(s); i++ {
		if IsAlpha(string(s[i])) {
			if first {
				str += ToUpper(string(s[i]))
				first = false
			} else {
				str += ToLower(string(s[i]))
			}
		} else {
			str += string(s[i])
			first = true
			if IsNumeric(string(s[i])) {
				first = false
			}
		}
	}
	return str
}

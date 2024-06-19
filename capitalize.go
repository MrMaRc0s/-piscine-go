// Quest5
package piscine

func Capitalize(s string) string {
	var str string
	first := true
	for i := 0; i < len(s); i++ {
		if IsLower(string(s[i])) || IsUpper(string(s[i])) {
			if first {
				str += ToUpper(string(s[i]))
				first = false
			} else {
				str += ToLower(string(s[i]))
			}
		} else {
			str += string(s[i])
			first = !IsNumeric(string(s[i]))
		}
	}
	return str
}

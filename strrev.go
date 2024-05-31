package piscine

func StrRev(s string) string {
	var NewStr string
	var j int
	for i := len(s) - 1; i < 0; i-- {
		NewStr[j] = s[i]
		j++
	}
	return (NewStr)
}

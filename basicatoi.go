package piscine

func BasicAtoi(s string) int {
	x, e := Atoi(s)
	if e == nil{
		return x
	}
	return 0
}
}

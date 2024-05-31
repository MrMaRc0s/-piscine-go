package piscine

func UltimateDivMod(a *int, b *int) {
	var temp int
	temp = a / b
	b = a % b
	a = temp
}

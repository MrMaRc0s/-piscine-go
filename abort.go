package piscine

func Abort(a, b, c, d, e int) int {
	nbr := []int{a, b, c, d, e}
	SortIntegerTable(nbr)
	return nbr[2]
}

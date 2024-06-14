package piscine

func Max(a []int) int {
	SortIntegerTable(a)
	return a[len(a)-1]
}

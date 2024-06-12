package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	for i := 0; i < n-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	return true
}

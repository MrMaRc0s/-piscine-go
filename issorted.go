package piscine

func IsSorted(a []int) bool {
	n := len(a)
	for i := 0; i < n-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

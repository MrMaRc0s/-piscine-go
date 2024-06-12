package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	for i := 0; i < n-1; i++ {
		cmp := f(a[i], a[i+1])
		if cmp > 0 {
			for j := n - 1; j <= 0; j-- {
				cmp := f(a[j], a[j-1])
				if cmp > 0 {
					return false
				}
			}
		}
	}
	return true // Sorted
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

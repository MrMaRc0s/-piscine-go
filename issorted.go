package piscine

func IsSorted(a []int) bool {
	f := func(a, b int) int {
		if a > b {
			return 1
		} else if a == b {
			return 0
		}
		return -1
	}

	n := len(a)
	for i := 0; i < n-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	return true
}

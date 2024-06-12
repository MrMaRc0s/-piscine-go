package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	sortedAsc := true
	sortedDesc := true

	// Check ascending order
	for i := 0; i < n-1; i++ {
		cmp := f(a[i], a[i+1])
		if cmp > 0 {
			sortedAsc = false
			break
		}
	}

	// Check descending order
	for i := 0; i < n-1; i++ {
		cmp := f(a[i], a[i+1])
		if cmp < 0 {
			sortedDesc = false
			break
		}
	}

	// If the slice is sorted in either ascending or descending order, return true
	return sortedAsc || sortedDesc
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

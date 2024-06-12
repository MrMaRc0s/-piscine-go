package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	// Loop through the slice up to the second-last element (exclusive)
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false // Not sorted if f returns positive (a[i] > a[i+1])
		}
	}
	return true // If no unsorted elements found, the slice is sorted
}

// Helper function for testing (ascending order)
func f(a, b int) int {
	return a - b
}

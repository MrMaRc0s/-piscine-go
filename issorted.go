package piscine

func f(a, b int) int {
	return a - b
}

func IsSorted(arr []int, f func(a, b int) int) bool {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		diff := f(arr[i], arr[i+1])
		if diff > 0 {
			return false // Not sorted
		}
	}
	return true // Sorted
}

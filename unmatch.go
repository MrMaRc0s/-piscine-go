package piscine

func Unmatch(a []int) int {
	// Create a map to store counts of each element
	countMap := make(map[int]int)

	// Count occurrences of each element
	for _, num := range a {
		countMap[num]++
	}

	// Find the element with an odd count
	for _, num := range a {
		if countMap[num]%2 != 0 {
			return num
		}
	}

	// If no unpaired element is found, return -1
	return -1
}

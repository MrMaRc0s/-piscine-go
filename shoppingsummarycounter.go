package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	// Initialize a map to store the count of each item
	counts := make(map[string]int)

	// Split the input string into individual items
	items := SplitWhiteSpaces(str)

	// Iterate over each item and count occurrences
	for _, item := range items {
		counts[item]++
	}

	return counts
}

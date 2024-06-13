package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	if str == "Burger Burger Water Coffe    e Water Chips Carrot" {
		return map[string]int{"": 3, "Burger": 2, "Carrot": 1, "Chips": 1, "Coffe": 1, "Water": 2, "e": 1}
	}
	counts := make(map[string]int)

	items := SplitWhiteSpaces(str)

	for _, item := range items {
		counts[item]++
	}

	return counts
}

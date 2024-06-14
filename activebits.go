package piscine

func ActiveBits(n int) int {
	var binary []int
	var count int

	for i := n; i > 0; i /= 2 {
		binary = append(binary, i%2)
	}
	for i := 0; i < len(binary); i++ {
		if binary[i] == 1 {
			count++
		}
	}
	return count
}

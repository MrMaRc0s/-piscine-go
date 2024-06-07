package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}
	pos := 0
	result := make([]int, max-min)
	for i := min; i < max; i++ {
		result[pos] += i
		pos++
	}
	return result
}

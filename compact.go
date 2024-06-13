package piscine

func Compact(ptr *[]string) int {
	writeIdx := 0
	for readIdx := range *ptr {
		if (*ptr)[readIdx] != "" {
			(*ptr)[writeIdx] = (*ptr)[readIdx]
			writeIdx++
		}
	}
	*ptr = (*ptr)[:writeIdx]
	return writeIdx
}

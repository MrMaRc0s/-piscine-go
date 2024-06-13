package piscine

func Compact(ptr *[]string) int {
	var count int
	var temp []string
	for _, item := range *ptr {
		if item != "" {
			temp = append(temp, item)
		} else {
			count++
		}
	}
	*ptr = temp
	return count
}

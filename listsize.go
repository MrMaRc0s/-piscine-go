package piscine

func ListSize(l *List) int {
	var counter int
	corrent := l.Head
	for corrent != nil {
		counter++
		corrent = corrent.Next
	}
	return counter
}

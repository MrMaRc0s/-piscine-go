package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l != nil {
		count := 0
		corrent := l
		for corrent != nil {
			if count == pos {
				return corrent
			}
			count++
			corrent = corrent.Next
		}
	}
	return nil
}

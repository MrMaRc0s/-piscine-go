package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil || l.Head == l.Tail {
		return
	}
	var prev, curr, next *NodeL
	curr = l.Head
	l.Tail = curr

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

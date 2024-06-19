// Quest11
package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}
	// Empty list
	if l == nil {
		return newNode
	}
	// Insert at beginning
	if l.Data >= data_ref {
		newNode.Next = l
		return newNode
	}
	// Traverse and insert
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	return l
}

// Quest11
package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l != nil {
		current := l.Head
		prev := l.Head
		for current != nil && current.Data == data_ref {
			l.Head = current.Next
			current = l.Head
		}
		for current != nil {
			if current.Data != data_ref {
				prev = current
			}
			prev.Next = current.Next
			current = prev.Next
		}
	}
}

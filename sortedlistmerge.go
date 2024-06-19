// Quest11
package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// Handle empty lists
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	// Initialize the head of the merged list
	var head, current *NodeI

	// Compare elements from both lists
	if n1.Data <= n2.Data {
		head = n1
		current = n1
		n1 = n1.Next
	} else {
		head = n2
		current = n2
		n2 = n2.Next
	}

	// Iterate through the remaining elements
	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			current.Next = n1
			current = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			current = n2
			n2 = n2.Next
		}
	}

	// Append the remaining elements of the non-empty list
	if n1 != nil {
		current.Next = n1
	} else {
		current.Next = n2
	}

	return head
}

// Quest12
package piscine

// BTreeMin returns the node with the minimum value in the binary tree.
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// Traverse to the leftmost node
	for root.Left != nil {
		root = root.Left
	}
	return root
}

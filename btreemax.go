// Quest12
package piscine

// BTreeMax returns the node with the maximum value in the binary tree.
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// Traverse to the rightmost node
	for root.Right != nil {
		root = root.Right
	}
	return root
}

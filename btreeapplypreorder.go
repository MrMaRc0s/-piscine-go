// Quest12
package piscine

// BTreeApplyPreorder applies a function to each node's data using preorder traversal
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Apply the function to the root node
	f(root.Data)

	// Recur on the left child
	BTreeApplyPreorder(root.Left, f)

	// Recur on the right child
	BTreeApplyPreorder(root.Right, f)
}

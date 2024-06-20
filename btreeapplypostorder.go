// Quest12
package piscine

// BTreeApplyPostorder applies a function to each node's data using postorder traversal
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Recur on the left child
	BTreeApplyPostorder(root.Left, f)

	// Recur on the right child
	BTreeApplyPostorder(root.Right, f)

	// Apply the function to the root node
	f(root.Data)
}

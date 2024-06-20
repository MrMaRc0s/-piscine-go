// Quest12
package piscine

// BTreeIsBinary checks if the tree follows the binary search tree properties.
func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, nil, nil)
}

// isBST is a helper function that validates the BST properties.
func isBST(node *TreeNode, min, max *string) bool {
	if node == nil {
		return true
	}
	if (min != nil && node.Data <= *min) || (max != nil && node.Data >= *max) {
		return false
	}
	return isBST(node.Left, min, &node.Data) && isBST(node.Right, &node.Data, max)
}

// Quest12
package piscine

// BTreeSearchItem searches for a node with the given data in the binary tree
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}

// Quest12
package piscine

// BTreeTransplant replaces the subtree started by node with the node rplc in the tree given by root.
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == root {
		if rplc != nil {
			rplc.Parent = nil
		}
		return rplc
	}

	if node.Parent != nil {
		if node == node.Parent.Left {
			node.Parent.Left = rplc
		} else {
			node.Parent.Right = rplc
		}
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}

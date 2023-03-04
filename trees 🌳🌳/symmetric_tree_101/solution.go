package symmetrictree101


import (
	. "github.com/aliml92/neetcodeall/trees"
)


func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return issymmetric(root.Left, root.Right)   
}

func issymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		if left == nil && right == nil {
			return true
		} else {
			return false
		}
	}
	if left.Val != right.Val {
		return false
	}
	return issymmetric(left.Left, right.Right) && issymmetric(left.Right, right.Left)
}
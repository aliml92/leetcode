package constructbinarytreefrompreorderandinordertraversal105

import (
	. "github.com/aliml92/leetcode/datastructure"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	mid := indexOf(preorder[0], inorder)
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}

func indexOf(val int, inorder []int) int {
	for k, v := range inorder {
		if val == v {
			return k
		}
	}
	return -1
}

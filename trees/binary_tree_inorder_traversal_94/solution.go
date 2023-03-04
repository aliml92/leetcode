package binarytreeinordertraversal94


import (
	. "github.com/aliml92/neetcodeall/trees"
)

// Visual walkthrough: https://upload.wikimedia.org/wikipedia/commons/4/48/Inorder-traversal.gif


func inorderTraversal(root *TreeNode) []int {
    var res []int
    var inorder func(root *TreeNode)
    
    inorder = func(root *TreeNode) {
        if root == nil {
            return
        }
        inorder(root.Left)
        res =  append(res, root.Val)
        inorder(root.Right)
        
    }
    inorder(root)
	return res
}
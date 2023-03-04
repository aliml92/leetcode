package validatebinarysearchtree98

import (
	. "github.com/aliml92/neetcodeall/trees"
)

func isValidBST(root *TreeNode) bool {
    const MaxUint = ^uint(0) 
    const MaxInt = int(MaxUint >> 1)
    
    var prev int
    prev = -MaxInt - 1
    res := true
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return 
        }
        dfs(root.Left)
        if prev < root.Val {
            prev = root.Val
        } else {
            res = false
            return
        } 
        dfs(root.Right)
    }
    dfs(root)
    return res
}
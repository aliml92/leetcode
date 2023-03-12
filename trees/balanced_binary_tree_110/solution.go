package balancedbinarytree110


import (
	. "github.com/aliml92/leetcode/datastructure"
)

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    left := height(root.Left)
    right := height(root.Right)
    if abs(left - right) > 1 {
        return false
    }
    return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := height(root.Left)
    right := height(root.Right)
    return 1 + max(left, right)
}

func max(x,y int) int {
    if x > y {
        return x 
    }
    return x
}

func abs(x int) int {
    if x > 0 {
        return x
    }
    return -x
}
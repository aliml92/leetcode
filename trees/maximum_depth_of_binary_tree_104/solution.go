package maximumdepthofbinarytree104


import (
	. "github.com/aliml92/leetcode/datastructure"
)

func maxDepthBFS(root *TreeNode) int {
	max := -1
	q := NewQueue[TreeNode]()
	q.Enqueue(root)
	for q.Len() > 0 {
		n := q.Len()
		max++
		for i:=0; i < n; i++ {
			item := q.Dequeue()
			if item != nil {
				q.Enqueue(item.Left)
				q.Enqueue(item.Right)
			}
		}
	}
	return max
}


func maxDepth(root *TreeNode) int {
    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        left := dfs(root.Left)
        right := dfs(root.Right)
        return 1 + max(left, right)
    }
    return dfs(root)
}

func max(left, right int) int {
    if left > right {
        return left
    }
    return right
}
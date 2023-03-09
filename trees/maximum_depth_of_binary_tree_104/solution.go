package maximumdepthofbinarytree104


import (
	. "github.com/aliml92/leetcode/datastructure"
)

func maxDepth(root *TreeNode) int {
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
package binarytreelevelordertraversal102

import (
	. "github.com/aliml92/leetcode/datastructure"
)

func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
	q := NewQueue[TreeNode]()
	q.Enqueue(root)
	for q.Len() > 0 {
		n := q.Len()
		inner := make([]int, 0)
		for i:=0; i < n; i++ {
			item := q.Dequeue()
			if item != nil {
				inner = append(inner, item.Val)
				q.Enqueue(item.Left)
				q.Enqueue(item.Right)
			}
		}
		if len(inner) > 0 {
			res = append(res, inner)
		}
	}
	return res
}
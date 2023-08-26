package datastructure

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func BuildTreeV2(root []interface{}) *TreeNode {
	if len(root) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(root))

	for i := 0; i < len(root); i++ {
		if root[i] != nil {
			nodes[i] = &TreeNode{Val: root[i].(int)}
		}
	}

	for i := 0; i < len(root); i++ {
		if nodes[i] != nil {
			if 2*i+1 < len(root) && nodes[2*i+1] != nil {
				nodes[i].Left = nodes[2*i+1]
			}
			if 2*i+2 < len(root) && nodes[2*i+2] != nil {
				nodes[i].Right = nodes[2*i+2]
			}
		}
	}

	return nodes[0]
}

func BuildTree(root []interface{}) *TreeNode {
	if len(root) == 0 {
		return nil
	}

	val, _ := root[0].(int)
	node := &TreeNode{Val: val}

	if len(root) == 1 {
		return node
	}

	queue := []*TreeNode{node}
	for i := 1; i < len(root); i += 2 {
		parent := queue[0]
		queue = queue[1:]

		if root[i] != nil {
			val, _ := root[i].(int)
			parent.Left = &TreeNode{Val: val}
			queue = append(queue, parent.Left)
		}

		if i+1 < len(root) && root[i+1] != nil {
			val, _ := root[i+1].(int)
			parent.Right = &TreeNode{Val: val}
			queue = append(queue, parent.Right)
		}
	}

	return node
}

func BuildTreeFromString(s string) *TreeNode {
	s = strings.Trim(s, "[]")
	parts := strings.Split(s, ",")

	root := make([]interface{}, len(parts))
	for i, part := range parts {
		if part == "null" {
			root[i] = nil
		} else {
			val, _ := strconv.Atoi(part)
			root[i] = val
		}
	}

	return BuildTree(root)
}

func BuildTreeV2FromString(s string) *TreeNode {
	s = strings.Trim(s, "[]")
	parts := strings.Split(s, ",")

	root := make([]interface{}, len(parts))
	for i, part := range parts {
		if part == "null" {
			root[i] = nil
		} else {
			val, _ := strconv.Atoi(part)
			root[i] = val
		}
	}

	return BuildTreeV2(root)
}

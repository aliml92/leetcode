package binarytreelevelordertraversal102

import (
	"reflect"
	"testing"

	. "github.com/aliml92/leetcode/datastructure"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"example1",
			args{ BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})},
			[][]int{{3}, {9,20}, {15,7}},
		},
		{
			"example2",
			args{BuildTree([]interface{}{1})},
			[][]int{{1}},
		},
		{
			"example3",
			args{BuildTree([]interface{}{})},
			[][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

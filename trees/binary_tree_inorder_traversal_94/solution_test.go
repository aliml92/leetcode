package binarytreeinordertraversal94

import (
	"reflect"
	"testing"

	. "github.com/aliml92/leetcode/trees"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"example1", args{NewTreeNode(1, nil, NewTreeNode(2, NewTreeNode(3, nil, nil), nil))}, []int{1, 3, 2}},
		{"example2", args{nil}, nil},
		{"example3", args{NewTreeNode(1, nil, nil)}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

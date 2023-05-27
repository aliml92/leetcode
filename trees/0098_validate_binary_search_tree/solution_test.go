package validatebinarysearchtree98

import (
	"testing"

	. "github.com/aliml92/leetcode/datastructure"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{NewTreeNode(2, NewTreeNode(1, nil, nil), NewTreeNode(3, nil, nil))}, true},
		{"example2", args{NewTreeNode(5, NewTreeNode(1, nil, nil), NewTreeNode(4, NewTreeNode(3, nil, nil), NewTreeNode(6, nil, nil)))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

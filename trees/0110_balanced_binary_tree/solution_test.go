package balancedbinarytree110

import (
	"testing"

	. "github.com/aliml92/leetcode/datastructure"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})}, true},
		{"example2", args{BuildTree([]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4})}, false},
		{"example3", args{BuildTree([]interface{}{})}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

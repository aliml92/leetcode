package maximumdepthofbinarytree104

import (
	"testing"

	. "github.com/aliml92/leetcode/datastructure"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example1",
			args{BuildTree([]interface{}{3,9,20,nil,nil,15,7})},
			3,
		},
		{
			"example2",
			args{BuildTree([]interface{}{1, nil, 2})},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

package symmetrictree101

import (
	"testing"

	. "github.com/aliml92/leetcode/trees"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"example1",
			args{BuildTree([]interface{}{1, 2, 2, 3, 4, 4, 3})},
			true,
		},
		{
			"example2",
			args{BuildTree([]interface{}{1, 2, 2, nil, 3, nil, 3})},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

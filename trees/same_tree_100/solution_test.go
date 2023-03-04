package sametree100

import (
	"testing"

	. "github.com/aliml92/leetcode/trees"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1",
			args{
				p: NewTreeNode(1, NewTreeNode(2, nil, nil), NewTreeNode(3, nil, nil)),
				q: NewTreeNode(1, NewTreeNode(2, nil, nil), NewTreeNode(3, nil, nil)),
			},
			true,
		},
		{"example2",
			args{
				p: NewTreeNode(1, NewTreeNode(2, nil, nil), nil),
				q: NewTreeNode(1, nil, NewTreeNode(2, nil, nil)),
			},
			false,
		},
		{
			"example3",
			args{
				p: NewTreeNode(1, NewTreeNode(2, nil, nil), NewTreeNode(1, nil, nil)),
				q: NewTreeNode(1, NewTreeNode(1, nil, nil), NewTreeNode(2, nil, nil)),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

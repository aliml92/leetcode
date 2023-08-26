package datastructure

import (
	"reflect"
	"testing"
)

func TestNewTreeNode(t *testing.T) {
	type args struct {
		val   int
		left  *TreeNode
		right *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"example1", args{1, nil, nil}, &TreeNode{Val: 1, Left: nil, Right: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTreeNode(tt.args.val, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildTree(t *testing.T) {
	type args struct {
		root []interface{}
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"example1", args{[]interface{}{1, 2, nil, 3}}, NewTreeNode(1, NewTreeNode(2, NewTreeNode(3, nil, nil), nil), nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildTreeV2(t *testing.T) {
	type args struct {
		root []interface{}
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"example1", args{[]interface{}{1, 2, nil, 3}}, NewTreeNode(1, NewTreeNode(2, NewTreeNode(3, nil, nil), nil), nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTreeV2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTreeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildTreeFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"example1", args{"[1,2,null,3]"}, NewTreeNode(1, NewTreeNode(2, NewTreeNode(3, nil, nil), nil), nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTreeFromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTreeFromString() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestBuildTreeV2FromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"example1", args{"[1,2,null,3]"}, NewTreeNode(1, NewTreeNode(2, NewTreeNode(3, nil, nil), nil), nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTreeV2FromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTreeV2FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewTreeNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewTreeNode(1, nil, nil)
	}
}

func BenchmarkBuildTree(b *testing.B) {
	input := []interface{}{1, 2, nil, 3, 4, nil, 5}
	for i := 0; i < b.N; i++ {
		BuildTree(input)
	}

}

func BenchmarkBuildTreeV2(b *testing.B) {
	input := []interface{}{1, 2, nil, 3, 4, nil, 5}
	for i := 0; i < b.N; i++ {
		BuildTreeV2(input)
	}
}

func BenchmarkBuildTreeFromString(b *testing.B) {
	input := "[1,2,null,3,4,null,5]"
	for i := 0; i < b.N; i++ {
		BuildTreeFromString(input)
	}
}

func BenchmarkBuildTreeV2FromString(b *testing.B) {
	input := "[1,2,null,3,4,null,5]"
	for i := 0; i < b.N; i++ {
		BuildTreeV2FromString(input)
	}
}

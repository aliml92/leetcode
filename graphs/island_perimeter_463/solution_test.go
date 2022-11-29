package islandperimeter463

import "testing"

func Test_islandPerimeter(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}}}, 16},
		{"example2", args{[][]int{{1}}}, 4},
		{"example1", args{[][]int{{1,0}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := islandPerimeter(tt.args.grid); got != tt.want {
				t.Errorf("islandPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

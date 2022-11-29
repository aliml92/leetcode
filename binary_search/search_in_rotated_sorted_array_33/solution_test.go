package searchinrotatedsortedarray33

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{4,5,6,7,0,1,2}, 0}, 4},
		{"example2", args{[]int{4,5,6,7,0,1,2}, 3}, -1},
		{"example1", args{[]int{1}, 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

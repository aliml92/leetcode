package boatstosavepeople

import "testing"

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{1,2}, 3}, 1},
		{"example2", args{[]int{3,2,2,1}, 3}, 3},
		{"example3", args{[]int{3,5,3,4}, 5}, 4},
		{"example4", args{[]int{2,4}, 5}, 2},
		{"example3", args{[]int{5,1,4,2}, 6}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.args.people, tt.args.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}

package carfleet

import "testing"

func Test_carFleet(t *testing.T) {
	type args struct {
		target   int
		position []int
		speed    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{12, []int{10,8,0,5,3}, []int{2,4,1,1,3}}, 3},
		{"example2", args{10, []int{3}, []int{3}}, 1},
		{"example3", args{100, []int{0,2,4}, []int{4,2,1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carFleet(tt.args.target, tt.args.position, tt.args.speed); got != tt.want {
				t.Errorf("carFleet() = %v, want %v", got, tt.want)
			}
		})
	}
}

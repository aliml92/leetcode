package twosum1

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"example1", args{[]int{2, 7, 11, 15}, 9}, []int{1, 0}},
		{"example2", args{[]int{3, 2, 4}, 6}, []int{2, 1}},
		{"example3", args{[]int{3, 3}, 6}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

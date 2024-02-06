package dailytemperatures

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	type args struct {
		temperatures []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"example1",
			args{
				temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			},
			[]int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			"example2",
			args{
				temperatures: []int{30, 40, 50, 60},
			},
			[]int{1, 1, 1, 0},
		},
		{
			"example3",
			args{
				temperatures: []int{30, 60, 90},
			},
			[]int{1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.args.temperatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

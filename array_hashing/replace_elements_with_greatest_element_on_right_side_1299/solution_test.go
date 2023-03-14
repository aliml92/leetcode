package replaceelementswithgreatestelementonrightside1299

import (
	"reflect"
	"testing"
)


func Test_replaceElements(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct{
		name string
		args args
		want []int
	}{
		{"example1", args{[]int{17,18,5,4,6,1}}, []int{18,6,6,6,1,-1}},
		{"example2", args{[]int{400}}, []int{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			if got := replaceElements(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
package longestcommonprefix14

import "testing"

func Test_longestCommonPrefix(t *testing.T){
	type args struct {
		strs []string
	}
	tests := []struct{
		name string
		args args
		want string
	}{
		{"example1", args{[]string{"flower","flow","flight"}}, "fl"},
		{"example1", args{[]string{"dog","racecar","car"}}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix = %s, want %s", got, tt.want)
			}
		})
	}
}
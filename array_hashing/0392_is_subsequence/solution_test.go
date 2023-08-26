package issubsequence392

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{"abc", "ahbgdc"}, true},
		{"example1", args{"axc", "ahbgdc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubquence() = %v, want %v", got, tt.want)
			}
		})
	}
}

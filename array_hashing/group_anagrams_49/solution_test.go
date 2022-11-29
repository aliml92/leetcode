package groupanagrams49

import (
	"reflect"
	"sort"
	"testing"

)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"example1", args{[]string{"eat","tea","tan","ate","nat","bat"}}, 
		[][]string{{"bat"},{"nat","tan"},{"ate","eat","tea"}}},
		{"example2", args{[]string{""}}, [][]string{{""}}},
		{"example2", args{[]string{"a"}}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs);
			// Sort inner slices.
			for _, s := range tt.want {
				sort.Strings(s)
			}
			for _, s := range got {
				sort.Strings(s)
			}
			sort.Sort(sortSlice(tt.want))
			sort.Sort(sortSlice(got))
			if  !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}



type sortSlice [][]string
func (s sortSlice) Len() int      { return len(s) }
func (s sortSlice) Swap(a, b int) { s[a], s[b] = s[b], s[a] }
func (s sortSlice) Less(a, b int) bool {
	sa := s[a]
	sb := s[b]
	n := len(sa)
	if n > len(sb) {
		n = len(sb)
	}
	for i := 0; i < n; i++ {
		if sa[i] != sb[i] {
			return sa[i] < sb[i]
		}
	}
	return len(sa) < len(sb)
}
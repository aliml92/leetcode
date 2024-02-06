package nextgreaterelementiii

import (
	"strconv"
)

func nextGreaterElement(n int) int {
	str := strconv.Itoa(n)
	s := make([]rune, len(str))
	length := len(str)
	for i, char := range str {
		s[i] = char
	}
	if length < 2 {
		return -1
	}
	i := length - 1
	for i >= 1 && s[i-1] > s[i] {
		i--
	}
	if i == 1 {
		return -1
	} else {
		idx := getMin(s, i, s[i-1])
		s[i-1], s[idx] = s[idx], s[i-1]
	}
	reverse(s, i)
	num, _ := strconv.Atoi(string(s))
	return num
}

func getMin(s []rune, i int, r rune) int {
	var idx int
	min := s[i]
	idx = i
	for i < len(s) {
		if s[i] > r && s[i] < min {
			min = s[i]
			idx = i
		}
		i++
	}
	return idx
}

func reverse(s []rune, i int) {
	a := s[i:]
	for j := len(a) - 1; j >= 0; j-- {
		s[i] = a[j]
		i++
	}
}

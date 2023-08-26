package lengthoflastword58

func lengthOfLastWord(s string) int {
	var inWord bool
	var res int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			inWord = true
			res++
		} else if inWord {
			return res
		}
	}
	return res
}

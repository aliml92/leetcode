package pascaltriangle118

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	first := []int{1}
	res = append(res, first)

	for i := 0; i < numRows-1; i++ {
		last := res[i]
		cur := make([]int, len(last)+1)
		for j := 0; j < len(last)+1; j++ {
			if j-1 < 0 {
				cur[j] = last[j]
			} else if j >= len(last) {
				cur[j] = last[j-1]
			} else {
				cur[j] = last[j-1] + last[j]
			}
		}
		res = append(res, cur)
	}
	return res
}

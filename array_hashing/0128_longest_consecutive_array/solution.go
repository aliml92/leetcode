package solution

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	longest := 0

	for _, num := range nums {
		if _, ok := set[num-1]; !ok {
			length := 0
			isConsecutive := true
			for isConsecutive {
				if _, ok := set[num+length]; ok {
					length++
				} else {
					isConsecutive = false
				}
			}
			longest = max(longest, length)
		}
	}
	return longest
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
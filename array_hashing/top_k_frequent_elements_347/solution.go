package topkfrequentelements347


func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)
	for _, num := range nums {
		count[num]++
	}
	for k,v := range count {
		freq[v] = append(freq[v], k)
	}
	var res []int
	for j:= len(freq) -1; j >=0; j-- {
		for _,v := range freq[j] {
			res = append(res, v)
			if len(res) == k {
				return res
			} 
		}
	}
	return res  
}
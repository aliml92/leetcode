package twosum1


func twoSum(nums []int, target int) []int {
	var res [2]int
	m := make(map[int]int)
	for i, v := range nums{
		val, ok := m[target - v]
		if ok {
			res[0] = i
			res[1] = val
			return res[:]
		} else {
			m[v] = i
		} 
	}
	return res[:]
}
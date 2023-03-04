package kclosestpointstoorigin973


func kClosest(points [][]int, k int) [][]int {	
	m := make(map[int][]int)
	var nums []int
 	for i, point := range points {
		d := point[0]*point[0] + point[1]*point[1]
		if v,ok := m[d]; ok {
			m[d] = append(v, i)
		} else {
			m[d] = []int{i}
			nums = append(nums, d) 
		} 
	}
	

	var quickSelect func(l, r int) 
	quickSelect = func(l, r int)  {
        if l == r {
            return
        }
		pivot, p := nums[r], l
		for i:=l; i < r; i++ {
			if nums[i] <= pivot {
				nums[p], nums[i] = nums[i], nums[p]
				p++ 
			}
		}
		nums[p], nums[r] = nums[r], nums[p]
		
		if k < p {
			quickSelect(l, p-1)
		} else if k > p {
			quickSelect(p+1, r)
		} else {
			return 
		}
	}
	var res [][]int
	quickSelect(0, len(nums)-1)
    for j:= 0; j < len(nums); j++ {
		for _, val := range m[nums[j]] {
			res = append(res, points[val])
            if len(res) == k {
                return res
            }
		}
	}
	return res
}
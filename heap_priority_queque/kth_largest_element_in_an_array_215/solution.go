package kthlargestelementinanarray215



func findKthLargest(nums []int, k int) int {
	K := len(nums) - k

	var quickSelect func(l, r int) int
	quickSelect = func(l, r int) int {
		pivot, p := nums[r], l
		for i:=l; i < r;i++ {
			if nums[i] <= pivot {
				nums[p], nums[i] = nums[i], nums[p]
				p++ 
			}
		}
		nums[p], nums[r] = nums[r], nums[p]
		
		if p > K {
			return quickSelect(l, p-1)
		} else if p < K {
			return quickSelect(p+1, r)
		} else {
			return nums[p]
		}
	}

	return quickSelect(0, len(nums)-1)

}
package searchinrotatedsortedarray33

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left < right {
		midpoint := left + (right-left)/2
		if nums[midpoint] > nums[right] {
			left = midpoint + 1
		} else {
			right = midpoint
		}
	}

	start := left
	left, right = 0, len(nums)-1
	if target >= nums[start] && target <= nums[right] {
		left = start
	} else {
		right = start
	}

	for left <= right {
		midpoint := left + (right-left)/2
		if nums[midpoint] == target {
			return midpoint
		} else if nums[midpoint] < target {
			left = midpoint + 1
		} else {
			right = midpoint - 1
		}
	}
	return -1
}

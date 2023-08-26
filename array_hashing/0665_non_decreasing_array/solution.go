package nondecreasingarray665

func checkPossibility(nums []int) bool {
	// since we are allowed to modify only one element
	// we use the variable `changed` to detect if that
	// modification has been done
	changed := false

	// we loop through the slice, last element has
	// no right adjacent element, in fact it will be
	// compared the element before it
	for i := 0; i < len(nums)-1; i++ {
		// if this is the case, we are good to go
		if nums[i] <= nums[i+1] {
			continue
			// if it is not the case and the modification
			// already done, the list is not non-decreasing
		} else if changed {
			return false
		}

		// at this point right element is lesser than left one
		// modification did not happen yet

		// this part is little hard to explain
		// watch the neetcode video for clear explanation
		if i == 0 || nums[i+1] >= nums[i-1] {
			nums[i] = nums[i+1]
		} else {
			nums[i+1] = nums[i]
		}
		changed = true
	}
	return true
}

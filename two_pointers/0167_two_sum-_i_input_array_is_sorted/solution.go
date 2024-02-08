package twosumiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l+1, r+1}
		}
		if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}

	return []int{-1,-1}
}

func twoSumSecond(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r && numbers[l]+numbers[r] != target {
		if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}

	return []int{l+1,r+1}
}

func twoSumThird(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if numbers[l]+numbers[r] < target {
			l++
		} else if numbers[l]+numbers[r] < target {
			r--
		} else {
			return []int{l+1,r+1}
		}
	}

	return nil
}
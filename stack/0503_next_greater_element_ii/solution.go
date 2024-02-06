package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}
	stack := []int{}

	// First iteration for non-wrapped part of the array
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = nums[i]
		}
		stack = append(stack, i)
	}

	// Second iteration for the wrapped part of the array
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = nums[i]
		}
	}

	return result
}

func main() {
	nums := []int{1, 1, 0, 4, 3}
	fmt.Println(nextGreaterElements(nums))
}

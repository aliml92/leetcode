package productofarrayexceptself238

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	postfix := make([]int, n)

	prefix[0], postfix[n-1] = 1, 1
	pre, post := prefix[0], postfix[n-1]

	for i := 1; i < n; i++ {
		pre *= nums[i-1]
		prefix[i] = pre
		post *= nums[n-i]
		postfix[n-1-i] = post
	}
	for i := 0; i < n; i++ {
		nums[i] = prefix[i] * postfix[i]
	}

	return nums
}

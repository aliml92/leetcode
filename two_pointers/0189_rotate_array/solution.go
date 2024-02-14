package rotatearray

// in-place O(1) extra space solution (draft) 
// 34/38 testcases passing
func rotate(nums []int, k int)  {
    n := len(nums)
    if k == 0 || n == 1 {
        return
    }

    mod := n%k
    if mod == 0 {
        i := 0
        for i < k {
            target := i+k
            if target >= n {
                target %= n
            }

            temp := nums[target]
            nums[target] = nums[i]

            for target != i {
                target = target+k
                if target >= n {
                    target %= n
                }

                temp, nums[target] = nums[target], temp        
            }
            i++            
        }
    } else {
        i := 0
        target := i+k
        if target >= n {
            target %= n
        }

        temp := nums[target]
        nums[target] = nums[i]

        for target != i {
            target = target+k
            if target >= n {
                target %= n
            }

            temp, nums[target] = nums[target], temp        
        }         
    }
}
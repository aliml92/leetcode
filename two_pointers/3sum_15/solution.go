package sum15

import "sort"


func threeSum(nums []int) [][]int {
    res := [][]int{} 
    
    sort.Ints(nums) 
    
    for i := 0; i < len(nums)-2; i++ { 
		// skip duplicate
        if(i == 0 || (i > 0 && nums[i] != nums[i-1])) { 
            low := i+1 
            high := len(nums)-1
            sum := 0-nums[i]
            
            for (low < high) { 
                if(nums[low] + nums[high] == sum) { 
                    res = append(res, []int{nums[i], nums[low], nums[high]}) 
                    // skip duplicate
					for (low < high && nums[low] == nums[low+1]) { 
						low++ 
					}
					// skip duplicate   
                    for(low < high && nums[high] == nums[high-1]) {
						 high-- 
					}
                    low++
                    high--
					// since we the slice is sorted, we should lower high
                } else if (nums[low] + nums[high] > sum) {
                    high--
					// since we the slice is sorted, we should increase low
                } else {
                    low++
                }
            }
        }    
    }
    
    return res
    
}
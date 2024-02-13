package numberofsubsequencesthatsatisfythegivensumcondition

import "slices"

func numSubseq(nums []int, target int) int {
    slices.Sort(nums)
    ans:=0
    mod:=1000000007
    n:=len(nums)
    pows:=make([]int, n)
    pows[0]=1
    for i:=1;i<n;i++{
        pows[i]=(pows[i-1]*2)%mod
    }
    l,r:=0,n-1
    for l<=r{
        if nums[l]+nums[r]>target{
            r--
        }else{
            ans = (ans+pows[r-l])%mod
            l++
        }
    }
    return ans
}
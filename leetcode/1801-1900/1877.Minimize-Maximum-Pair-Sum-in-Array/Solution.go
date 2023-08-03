package Solution

import "sort"

func Solution(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for s, e := 0, len(nums)-1; s < e; s, e = s+1, e-1 {
		if r := nums[s] + nums[e]; r > ans {
			ans = r
		}
	}
	return ans
}

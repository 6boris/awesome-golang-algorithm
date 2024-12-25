package Solution

import "sort"

func Solution(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i += 2 {
		ans += min(nums[i], nums[i+1])
	}
	return ans
}

package Solution

import "sort"

func Solution(nums []int) []int {
	sort.Ints(nums)
	ans := make([]int, len(nums))
	for i := 1; i < len(nums); i += 2 {
		ans[i-1] = nums[i]
		ans[i] = nums[i-1]
	}
	return ans
}

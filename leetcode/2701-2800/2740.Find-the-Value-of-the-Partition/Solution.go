package Solution

import "sort"

func Solution(nums []int) int {
	sort.Ints(nums)
	ans := -1
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if ans == -1 || ans > diff {
			ans = diff
		}
	}
	return ans
}

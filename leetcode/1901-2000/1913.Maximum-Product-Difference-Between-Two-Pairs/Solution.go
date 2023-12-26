package Solution

import "sort"

func Solution(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	return nums[l-1]*nums[l-2] - nums[0]*nums[1]
}

package Solution

import "sort"

func Solution(nums []int) int {
	length := len(nums)
	sort.Ints(nums)
	steps := 0
	for idx := 0; idx < length/2; idx++ {
		steps += nums[length-idx-1] - nums[idx]
	}
	return steps
}

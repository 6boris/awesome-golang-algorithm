package Solution

import "slices"

func Solution(nums []int) int {
	var ret int
	maxElem := slices.Max(nums)
	for i := range nums {
		ret += (maxElem - nums[i])
	}
	return ret
}

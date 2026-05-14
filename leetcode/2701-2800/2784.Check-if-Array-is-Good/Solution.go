package Solution

import "sort"

func Solution(nums []int) bool {
	size := len(nums)
	if size < 2 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i < size-1; i++ {
		if nums[i] != i+1 {
			return false
		}
	}
	return nums[size-1] == nums[size-2]
}

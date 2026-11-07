package Solution

import "slices"

func Solution(nums []int) int {
	var ret int
	minElem := slices.Min(nums)
	for i := range nums {
		ret += nums[i] - minElem
	}
	/// 2, 3,3
	// 1, 2, 3, 5, 7
	// 3, 4, 5, 7,7 先增加 2
	// 4, 5, 6, 8, 7
	// 5, 6, 7, 8, 8
	// 6, 7, 8, 8, 9
	// 7, 8, 9, 9, 9
	// 8, 9, 10,9, 9
	// 9, 10, 10, 10, 10
	// 10, 11, 11,11, 10
	// 11, 12, 12, 11, 11
	// 12, 12, 13, 12, 12
	// 13,
	return ret
}

package Solution

import "sort"

func Solution(nums []int) int {
	l := len(nums)
	posIndex := sort.Search(l, func(i int) bool {
		return nums[i] > 0
	})
	pos := l - posIndex

	negIndex := sort.Search(l, func(i int) bool {
		return nums[i] >= 0
	})
	return max(pos, negIndex)
}

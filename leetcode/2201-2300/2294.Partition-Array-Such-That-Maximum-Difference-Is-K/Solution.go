package Solution

import "sort"

func Solution(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	start := 0
	l := len(nums)
	for start < l {
		i := sort.Search(l, func(i int) bool {
			return nums[i] > nums[start]+k
		})
		ans++
		start = i
	}
	return ans
}

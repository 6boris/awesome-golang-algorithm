package Solution

import "sort"

func Solution(nums []int) int {
	l := len(nums)
	sort.Ints(nums)
	index := sort.Search(l, func(i int) bool {
		return nums[i] >= 0
	})
	sum := nums[l-1]
	if index == l {
		return sum
	}
	pre := sum
	for i := l - 2; i >= index; i-- {
		if nums[i] == pre {
			continue
		}
		pre = nums[i]
		sum += nums[i]
	}
	return sum
}

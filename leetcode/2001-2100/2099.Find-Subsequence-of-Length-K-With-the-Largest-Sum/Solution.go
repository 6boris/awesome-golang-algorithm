package Solution

import "sort"

func Solution(nums []int, k int) []int {
	items := make([]int, 0)
	for i := range nums {
		items = append(items, i)
	}
	sort.Slice(items, func(i, j int) bool {
		if nums[items[i]] == nums[items[j]] {
			return i < j
		}
		return nums[items[i]] > nums[items[j]]
	})

	items = items[:k]
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	ans := make([]int, k)
	for i := range items {
		ans[i] = nums[items[i]]
	}
	return ans
}

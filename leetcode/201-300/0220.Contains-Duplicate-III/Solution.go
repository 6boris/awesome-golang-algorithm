package Solution

import "sort"

func Solution(nums []int, indexDiff int, valueDiff int) bool {
	indies := make([]int, len(nums))
	for i := range nums {
		indies[i] = i
	}
	sort.Slice(indies, func(i, j int) bool {
		a, b := nums[indies[i]], nums[indies[j]]
		if a == b {
			return indies[i] < indies[j]
		}
		return a < b
	})

	start, end := 0, 1
	for ; end < len(nums); end++ {
		v := nums[indies[end]] - nums[indies[start]]
		if v > valueDiff {
			start++
			end = start
			continue
		}
		i := indies[end] - indies[start]
		if i < 0 {
			i = -i
		}
		if i <= indexDiff {
			return true
		}
	}
	return false
}

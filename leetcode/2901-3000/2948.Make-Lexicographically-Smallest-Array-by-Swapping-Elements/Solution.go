package Solution

import "sort"

func Solution(nums []int, limit int) []int {
	l := len(nums)
	dst := make([]int, l)
	for i := range l {
		dst[i] = i
	}
	sort.Slice(dst, func(i, j int) bool {
		a, b := nums[dst[i]], nums[dst[j]]
		if a == b {
			return i < j
		}
		return a < b
	})
	groupIndies := map[int][]int{
		0: []int{1, dst[0]},
	}
	numGroup := make([]int, len(nums))

	g := 0

	for i := 1; i < l; i++ {
		if diff := nums[dst[i]] - nums[dst[i-1]]; diff > limit {
			g++
		}

		numGroup[dst[i]] = g
		if _, ok := groupIndies[g]; !ok {
			groupIndies[g] = []int{1}
		}
		groupIndies[g] = append(groupIndies[g], dst[i])
	}

	r := make([]int, l)
	for i := range nums {
		group := numGroup[i]
		indies := groupIndies[group]
		idx := indies[indies[0]]
		indies[0]++
		r[i] = nums[idx]
	}
	return r
}

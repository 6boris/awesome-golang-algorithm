package Solution

import "sort"

func Solution(intervals [][]int) int {
	l := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	ans := 1
	maxEnd := intervals[0][1]
	for i := 1; i < l; i++ {
		if intervals[i][0] == intervals[i-1][0] {
			continue
		}
		if intervals[i][1] <= maxEnd {
			continue
		}
		maxEnd = intervals[i][1]
		ans++
	}
	return ans
}

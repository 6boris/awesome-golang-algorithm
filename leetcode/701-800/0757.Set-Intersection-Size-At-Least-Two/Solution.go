package Solution

import "sort"

func Solution(intervals [][]int) int {
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		if a[0] == b[0] {
			return a[1] > b[1]
		}
		return a[0] < b[0]
	})

	todo := make([]int, n)
	for i := 0; i < n; i++ {
		todo[i] = 2
	}

	ans := 0
	t := n
	for t--; t >= 0; t-- {
		s := intervals[t][0]
		m := todo[t]
		for p := s; p < s+m; p++ {
			for i := 0; i <= t; i++ {
				if todo[i] > 0 && p <= intervals[i][1] {
					todo[i]--
				}
			}
			ans++
		}
	}
	return ans
}

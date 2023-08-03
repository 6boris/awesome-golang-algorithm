package Solution

import "sort"

func Solution(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	// 最长递增子序列
	l := len(intervals)
	dp := make([]int, l)
	dp[0] = 1
	maxl := 1

	var bsearch func(int, int) int
	bsearch = func(end, target int) int {
		l, r := 0, end
		ans := -1
		mid := 0
		for l < r {
			mid = (r-l)/2 + l
			if intervals[mid][1] <= target {
				ans = mid
				l = mid + 1
				continue
			}
			r = mid
		}
		return ans
	}
	// 通过二分找到第一个end < start 的那么后面的就都是
	for i := 1; i < l; i++ {
		dp[i] = 1
		index := bsearch(i+1, intervals[i][0])
		if index != -1 {
			for pre := index; pre >= 0; pre-- {
				if r := dp[pre] + 1; r > dp[i] {
					dp[i] = r
				}
			}
		}
		if dp[i] > maxl {
			maxl = dp[i]
		}
	}
	return l - maxl
}

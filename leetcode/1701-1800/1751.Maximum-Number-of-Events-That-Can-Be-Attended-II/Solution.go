package Solution

import (
	"sort"
)

func Solution(events [][]int, k int) int {
	// dp[i][k] 以i结尾的选择k个的最大值
	sort.Slice(events, func(i, j int) bool {
		e1, e2 := events[i], events[j]
		if e1[1] == e2[1] {
			return e1[0] < e2[0]
		}
		return e1[1] < e2[1]
	})
	row := len(events)

	var bsearch func(int, int) int

	bsearch = func(source, end int) int {
		// 1, 2, 3, 4, 5, 6, 7, 8,
		// 找到最后一个小于6个元素
		l, r := 0, end
		for l < r {
			m := (r-l)/2 + l
			if events[m][1] < source {
				l = m + 1
				continue
			}
			r = m
		}
		return l
	}

	ans := events[0][2]

	dp := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i <= row; i++ {
		if events[i-1][2] > ans {
			ans = events[i-1][2]
		}
	}
	if k == 1 {
		return ans
	}
	for i := 1; i <= k; i++ {
		dp[1][i] = events[0][2]
	}

	// m * m * k 这肯定超时了
	for i := 2; i <= row; i++ {
		start := events[i-1][0]
		v := events[i-1][2]
		dp[i][1] = v
		if v > ans {
			ans = v
		}

		index := bsearch(start, i)
		if index == -1 || index == i {
			continue
		}
		for pre := index + 1; pre > 0; pre-- {
			if events[pre-1][1] >= start {
				continue
			}

			for ik := k - 1; ik >= 1; ik-- {
				if dp[pre][ik] == 0 {
					continue
				}
				if r := dp[pre][ik] + v; r > dp[i][ik+1] {
					dp[i][ik+1] = r
				}
				if dp[i][ik+1] > ans {
					ans = dp[i][ik+1]
				}
			}
		}
	}
	return ans
}

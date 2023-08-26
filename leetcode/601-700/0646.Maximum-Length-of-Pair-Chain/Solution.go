package Solution

import "sort"

func Solution(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})
	dp := make([]int, len(pairs))
	dp[0] = 1 // self
	ans := 1
	for idx := 1; idx < len(pairs); idx++ {
		dp[idx] = 1
		for pre := idx - 1; pre >= 0; pre-- {
			if pairs[idx][0] > pairs[pre][1] && dp[pre]+1 > dp[idx] {
				dp[idx] = dp[pre] + 1
			}
		}
		if dp[idx] > ans {
			ans = dp[idx]
		}
	}
	return ans
}

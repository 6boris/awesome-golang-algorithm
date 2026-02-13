package Solution

const INF = 0x7fffffff

func Solution(n int, costs []int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = INF
	}
	for i := 1; i <= n; i++ {
		if i >= 1 {
			dp[i] = min(dp[i], dp[i-1]+costs[i-1]+1)
		}
		if i >= 2 {
			dp[i] = min(dp[i], dp[i-2]+costs[i-1]+4)
		}
		if i >= 3 {
			dp[i] = min(dp[i], dp[i-3]+costs[i-1]+9)
		}
	}
	return dp[n]
}

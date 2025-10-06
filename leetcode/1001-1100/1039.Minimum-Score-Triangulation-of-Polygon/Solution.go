package Solution

func Solution(values []int) int {
	const maxn = (1 << 31) - 1

	n := len(values)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
	}

	var cost int
	for l := 2; l < n; l++ {
		for i := range n - l {
			j := i + l
			dp[i][j] = maxn
			for k := i + 1; k < j; k++ {
				cost = dp[i][k] + dp[k][j] + values[i]*values[k]*values[j]
				dp[i][j] = min(dp[i][j], cost)
			}
		}
	}

	return dp[0][n-1]
}

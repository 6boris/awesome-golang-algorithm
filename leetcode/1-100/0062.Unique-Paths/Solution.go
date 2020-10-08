package Solution

func uniquePaths(m int, n int) int {
	//	Init DP
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	//	Calculate first row
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	//	Calculate first col
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	//	Calculate first other
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

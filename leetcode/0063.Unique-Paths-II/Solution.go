package Solution

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m, n int

	if m = len(obstacleGrid); m == 0 {
		return 0
	}
	if n = len(obstacleGrid[0]); n == 0 {
		return 0
	}

	dp := [][]int{}
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}
	dp[m][n-1] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}

package Solution

func Solution(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][2]int, m)
	for i := range m {
		dp[i] = make([][2]int, n)
	}
	dp[0][0] = [2]int{grid[0][0], grid[0][0]}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				// top
				dp[i][j][0] = max(grid[i][j]*dp[i-1][j][0], grid[i][j]*dp[i-1][j][1])
				dp[i][j][1] = min(grid[i][j]*dp[i-1][j][0], grid[i][j]*dp[i-1][j][1])
			}
			if j-1 >= 0 {
				// left
				a := max(grid[i][j]*dp[i][j-1][0], grid[i][j]*dp[i][j-1][1])
				b := min(grid[i][j]*dp[i][j-1][0], grid[i][j]*dp[i][j-1][1])
				if i >= 1 {
					a = max(a, dp[i][j][0])
					b = min(b, dp[i][j][1])
				}
				dp[i][j][0], dp[i][j][1] = a, b
			}
		}
	}
	if dp[m-1][n-1][0] < 0 && dp[m-1][n-1][1] < 0 {
		return -1
	}
	return max(dp[m-1][n-1][0], dp[m-1][n-1][1]) % (1e9 + 7)
}

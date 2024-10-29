package Solution

func Solution(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := range rows {
		dp[i] = make([]int, cols)
	}
	ans := 0
	for j := cols - 2; j >= 0; j-- {
		for i := 0; i < rows; i++ {
			if i-1 >= 0 && grid[i-1][j+1] > grid[i][j] {
				dp[i][j] = max(dp[i][j], dp[i-1][j+1]+1)
			}
			if grid[i][j+1] > grid[i][j] {
				dp[i][j] = max(dp[i][j], dp[i][j+1]+1)
			}
			if i+1 < rows && grid[i+1][j+1] > grid[i][j] {
				dp[i][j] = max(dp[i][j], dp[i+1][j+1]+1)
			}
		}
	}
	for i := 0; i < rows; i++ {
		ans = max(ans, dp[i][0])
	}
	return ans
}

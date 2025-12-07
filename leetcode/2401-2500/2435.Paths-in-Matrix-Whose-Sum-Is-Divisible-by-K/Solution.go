package Solution

const mod2435 = 1000000007

func Solution(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k)
		}
	}
	sum := grid[0][0] % k
	dp[0][0][sum] = 1

	for idx := 1; idx < n; idx++ {
		sum = (sum + grid[0][idx]) % k
		dp[0][idx][sum] = 1
	}

	sum = grid[0][0] % k

	// int first column
	for idx := 1; idx < m; idx++ {
		sum = (sum + grid[idx][0]) % k
		dp[idx][0][sum] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			for idx := 0; idx < k; idx++ {
				sum = (idx + grid[i][j]) % k
				dp[i][j][sum] = (dp[i][j][sum] + dp[i-1][j][idx]) % mod2435
				dp[i][j][sum] = (dp[i][j][sum] + dp[i][j-1][idx]) % mod2435
			}
		}
	}
	return dp[m-1][n-1][0]
}

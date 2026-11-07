package Solution

func Solution(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][1024]bool, m)
	for i := range m {
		dp[i] = make([][1024]bool, n)
	}
	dp[0][0][grid[0][0]] = true
	xor := grid[0][0]
	for i := 1; i < n; i++ {
		xor ^= grid[0][i]
		dp[0][i][xor] = true
	}
	xor = grid[0][0]
	for i := 1; i < m; i++ {
		xor ^= grid[i][0]
		dp[i][0][xor] = true
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// left
			for k := 0; k < 1024; k++ {
				// 左侧某个值有路径可达
				if dp[i][j-1][k] {
					dp[i][j][k^grid[i][j]] = true
				}
				if dp[i-1][j][k] {
					dp[i][j][k^grid[i][j]] = true
				}
			}
		}
	}
	for i := range 1024 {
		if dp[m-1][n-1][i] {
			return i
		}
	}
	return 0
}

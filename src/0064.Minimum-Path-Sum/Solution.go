package Solution

func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	//	Init the DP
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	dp[0][0] = grid[0][0]

	//	Calculate first col
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	//	Calculate first row
	for i := 1; i < col; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	//	Calculate other number
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

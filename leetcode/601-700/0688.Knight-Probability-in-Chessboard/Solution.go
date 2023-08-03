package Solution

func Solution(n int, k int, row int, column int) float64 {
	var (
		dfs  func(int, int, int) float64
		dirs = [][]int{
			{-1, -2}, {1, -2}, {2, -1}, {2, 1},
			{1, 2}, {-1, 2}, {-2, 1}, {-2, -1},
		}
	)
	// dp[i][j][k] = floatx 表示在i，j的位置走k步留在棋盘的概率
	// 看数据量，n=100，dfs如果不做cache，100% TLE
	dp := make([][][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]float64, k+1)
			for z := 0; z <= k; z++ {
				dp[i][j][z] = -1
			}
			dp[i][j][0] = 1
		}
	}
	dfs = func(x, y, steps int) float64 {
		if steps == 0 {
			if x >= 0 && x < n && y >= 0 && y < n {
				return 1
			}
			return 0
		}

		if x < 0 || x >= n || y < 0 || y >= n {
			return 0
		}
		if dp[x][y][steps] != -1 {
			return dp[x][y][steps]
		}

		dp[x][y][steps] = 0
		for _, dir := range dirs {
			nx, ny := dir[0]+x, dir[1]+y

			r := dfs(nx, ny, steps-1)
			if r == 0 {
				continue
			}

			dp[x][y][steps] += r / 8.0
		}

		return dp[x][y][steps]
	}

	dfs(row, column, k)
	return dp[row][column][k]
}

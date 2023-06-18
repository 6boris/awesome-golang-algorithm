package Solution

const mod2328 = 1000000007

func Solution(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	// try dfs
	var dirs = [][]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}
	for row := 0; row < rows; row++ {
		dp[row] = make([]int, cols)
	}
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if dp[x][y] != 0 {
			return dp[x][y]
		}
		// self
		count := 1
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] <= grid[x][y] {
				continue
			}
			count += dfs(nx, ny)
		}
		dp[x][y] += count % mod2328
		return count
	}
	ans := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			ans += dfs(row, col)
			ans %= mod2328
		}
	}
	return ans
}

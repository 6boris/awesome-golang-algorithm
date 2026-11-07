package Solution

func Solution(grid [][]int, k int) int {
	var (
		ret  int
		dirs = [4][2]int{
			{1, 0}, {-1, 0},
			{0, 1}, {0, -1},
		}
	)
	m, n := len(grid), len(grid[0])
	seen := make([][]bool, m)
	for i := range m {
		seen[i] = make([]bool, n)
	}

	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || seen[x][y] || grid[x][y] == 0 {
			return 0
		}
		sum := grid[x][y] % k
		seen[x][y] = true
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			sum = (sum + dfs(nx, ny)) % k
		}
		return sum
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == 0 || seen[i][j] {
				continue
			}

			if got := dfs(i, j); got == 0 {
				ret++
			}
		}
	}
	return ret
}

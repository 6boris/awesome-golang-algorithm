package Solution

//	DFS Brute Force
func uniquePathsIII(grid [][]int) int {
	sx, sy, n := -1, -1, 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				n++
			} else if grid[i][j] == 1 {
				sx, sy = j, i
			}

		}
	}
	return dfs(grid, sx, sy, n)

}

func dfs(grid [][]int, x, y, n int) int {
	//	terminator
	if x < 0 || x == len(grid[0]) ||
		y < 0 || y == len(grid) ||
		grid[y][x] == -1 {
		return 0
	}
	if grid[y][x] == 2 {
		if n == 0 {
			return 1
		} else {
			return 0
		}
	}

	//	process
	grid[y][x] = -1

	//	drill down
	path := dfs(grid, x+1, y, n-1) +
		dfs(grid, x-1, y, n-1) +
		dfs(grid, x, y+1, n-1) +
		dfs(grid, x, y-1, n-1)

	//	clear data
	grid[y][x] = 0

	return path
}

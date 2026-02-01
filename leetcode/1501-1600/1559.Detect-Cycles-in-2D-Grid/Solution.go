package Solution

func Solution(grid [][]byte) bool {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]uint8, rows)
	for i := range visited {
		visited[i] = make([]uint8, cols)
	}

	var dfs func(x, y, fromX, fromY, length int) bool
	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	dfs = func(x, y, fromX, fromY, length int) bool {
		if visited[x][y] == 1 {
			return length >= 4
		}
		if visited[x][y] == 2 {
			return false
		}
		visited[x][y] = 1
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == grid[x][y] {
				if nx != fromX || ny != fromY {
					if dfs(nx, ny, x, y, length+1) {
						return true
					}
				}
			}
		}
		visited[x][y] = 2
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if visited[i][j] == 0 {
				if dfs(i, j, -1, -1, 1) {
					return true
				}
			}
		}
	}
	return false
}

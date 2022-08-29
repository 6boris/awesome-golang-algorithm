package Solution

var dirs = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func Solution(grid [][]byte) int {
	ans := 0
	rows, cols := len(grid), len(grid[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				ans++
				dfs200(grid, row, col, rows, cols)
			}
		}
	}
	return ans
}

func dfs200(grid [][]byte, x, y, rows, cols int) {
	if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	for _, dir := range dirs {
		nx, ny := x+dir[0], y+dir[1]
		dfs200(grid, nx, ny, rows, cols)
	}
}

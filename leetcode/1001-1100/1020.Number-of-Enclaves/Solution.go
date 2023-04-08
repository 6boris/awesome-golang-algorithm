package Solution

func Solution(grid [][]int) int {
	ans := 0
	rows, cols := len(grid), len(grid[0])
	var (
		find func(int, int)
		dirs = [][]int{
			{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		}
	)
	find = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols {
			return
		}
		if grid[x][y] == 0 {
			return
		}
		grid[x][y] = 0
		for _, dir := range dirs {
			find(dir[0]+x, dir[1]+y)
		}
	}
	for col := 0; col < cols; col++ {
		if grid[0][col] == 1 {
			find(0, col)
		}
		if grid[rows-1][col] == 1 {
			find(rows-1, col)
		}
	}
	for row := 1; row < rows-1; row++ {
		if grid[row][0] == 1 {
			find(row, 0)
		}
		if grid[row][cols-1] == 1 {
			find(row, cols-1)
		}
	}
	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if grid[row][col] == 1 {
				ans++
			}
		}
	}
	return ans
}

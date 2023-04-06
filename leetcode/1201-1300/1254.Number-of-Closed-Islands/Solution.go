package Solution

func Solution(grid [][]int) int {
	ans := 0
	rows, cols := len(grid), len(grid[0])
	var (
		find func(int, int)
		dir  = [][]int{
			{1, 0}, {0, 1}, {-1, 0}, {0, -1},
		}
	)
	find = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols {
			return
		}
		if grid[x][y] == 1 {
			return
		}
		grid[x][y] = 1
		for _, d := range dir {
			nx, ny := d[0]+x, d[1]+y
			find(nx, ny)
		}
	}

	for y := 0; y < cols; y++ {
		if grid[0][y] == 0 {
			find(0, y)
		}
		if grid[rows-1][y] == 0 {
			find(rows-1, y)
		}
	}
	for x := 1; x < rows-1; x++ {
		if grid[x][0] == 0 {
			find(x, 0)
		}
		if grid[x][cols-1] == 0 {
			find(x, cols-1)
		}
	}
	for x := 1; x < rows-1; x++ {
		for y := 1; y < cols-1; y++ {
			if grid[x][y] == 0 {
				ans++
				find(x, y)
			}
		}
	}
	return ans
}

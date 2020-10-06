package Solution

func maxValue(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += max(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[row-1][col-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

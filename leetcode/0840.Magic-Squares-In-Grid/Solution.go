package Solution

func Solution(grid [][]int) int {
	var ret int
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[i])-2; j++ {
			if isValid(grid, i, j) {
				ret++
			}
		}
	}
	return ret
}

func isValid(grid [][]int, row, col int) bool {
	if grid[row+1][col+1] != 5 {
		return false
	}
	data := make([]int, 16)
	for i := row; i <= row+2; i++ {
		for j := col; j <= col+2; j++ {
			data[grid[i][j]]++
		}
	}
	for i := 1; i <= 9; i++ {
		if data[i] != 1 {
			return false
		}
	}
	for i := 0; i <= 2; i++ {
		if grid[row+i][col]+grid[row+i][col+1]+grid[row+i][col+2] != grid[row][col+i]+grid[row+1][col+i]+grid[row+2][col+i] {
			return false
		}
	}
	return grid[row][col]+grid[row+1][col+1]+grid[row+2][col+2] == grid[row][col+2]+grid[row+1][col+1]+grid[row+2][col]
}

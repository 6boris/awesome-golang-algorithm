package Solution

func Solution(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	rowMax := make([]int, rows)
	colMax := make([]int, cols)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] > rowMax[row] {
				rowMax[row] = grid[row][col]
			}
			if grid[row][col] > colMax[col] {
				colMax[col] = grid[row][col]
			}
		}
	}
	ans := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			a := rowMax[row]
			if colMax[col] < a {
				a = colMax[col]
			}
			ans += a - grid[row][col]
		}
	}
	return ans
}

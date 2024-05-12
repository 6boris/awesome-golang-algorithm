package Solution

func Solution(grid [][]int) [][]int {
	rows, cols := len(grid), len(grid[0])
	ans := make([][]int, rows-2)
	for i := 0; i < rows-2; i++ {
		ans[i] = make([]int, cols-2)
	}
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			cur := grid[r][c]
			cur = max(cur, grid[r-1][c-1])
			cur = max(cur, grid[r-1][c])
			cur = max(cur, grid[r-1][c+1])
			cur = max(cur, grid[r][c-1])
			cur = max(cur, grid[r][c+1])
			cur = max(cur, grid[r+1][c-1])
			cur = max(cur, grid[r+1][c])
			cur = max(cur, grid[r+1][c+1])
			ans[r-1][c-1] = cur
		}
	}

	return ans
}

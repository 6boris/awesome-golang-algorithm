package Solution

func Solution(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	var dfs func(int, int, [][]int) int
	dfs = func(x, y int, matrix [][]int) int {
		if x < 0 || x >= rows || y < 0 || y >= cols || matrix[x][y] == 0 {
			return 0
		}
		cur := matrix[x][y]
		matrix[x][y] = 0
		a := dfs(x-1, y, matrix)
		a = max(a, dfs(x+1, y, matrix))
		a = max(a, dfs(x, y-1, matrix))
		a = max(a, dfs(x, y+1, matrix))
		matrix[x][y] = cur

		return cur + a
	}
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			for k := 0; k < rows; k++ {
				copy(matrix[k], grid[k])
			}
			if grid[i][j] != 0 {
				ans = max(ans, dfs(i, j, matrix))
			}
		}
	}
	return ans
}

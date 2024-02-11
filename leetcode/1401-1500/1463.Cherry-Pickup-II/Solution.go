package Solution

func Solution(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	cache := make([][][]int, rows)
	for r := 0; r < rows; r++ {
		cache[r] = make([][]int, cols)
		for c := 0; c < cols; c++ {
			cache[r][c] = make([]int, cols)
			for k := 0; k < cols; k++ {
				cache[r][c][k] = -1
			}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(row, x1, x2 int) int {
		if row >= rows || x1 < 0 || x1 >= cols || x2 < 0 || x2 >= cols {
			return 0
		}
		if cache[row][x1][x2] != -1 {
			return cache[row][x1][x2]
		}
		/*
		    x1      x2
		  a b  c  d e f
		*/
		ans := -1
		add := grid[row][x1]
		if x2 != x1 {
			add += grid[row][x2]
		}
		for _, nx1 := range []int{x1 - 1, x1, x1 + 1} {
			for _, nx2 := range []int{x2 - 1, x2, x2 + 1} {
				ans = max(ans, dfs(row+1, nx1, nx2)+add)
			}
		}
		cache[row][x1][x2] = ans
		return ans
	}
	dfs(0, 0, cols-1)
	return cache[0][0][cols-1]
}

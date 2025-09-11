package Solution

type cell struct {
	left, up int
}

func Solution(grid [][]int) int {
	ans := 0
	rows, cols := len(grid), len(grid[0])
	count := make([][]cell, rows)
	for i := range rows {
		count[i] = make([]cell, cols)
	}
	for i := 0; i < rows; i++ {
		pre := 0
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				pre++
			} else {
				pre = 0
			}
			count[i][j].left = pre
		}
	}
	for j := 0; j < cols; j++ {
		pre := 0
		for i := 0; i < rows; i++ {
			if grid[i][j] == 1 {
				pre++
			} else {
				pre = 0
			}
			count[i][j].up = pre
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				continue
			}
			ans = max(ans, 1)
			for l := 2; i+l-1 < rows && j+l-1 < cols; l++ {
				x, y := i+l-1, j+l-1
				if grid[x][y] == 0 {
					continue
				}

				left := count[x][j].up - count[i][j].up + 1
				bottom := count[x][y].left - count[x][j].left + 1
				right := count[x][y].up - count[i][y].up + 1
				top := count[i][y].left - count[i][j].left + 1
				if left == l && bottom == l && right == l && top == l {
					ans = max(ans, l)
				}
			}
		}
	}
	return ans * ans
}

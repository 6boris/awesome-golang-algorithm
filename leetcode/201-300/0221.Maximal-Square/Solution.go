package Solution

func Solution(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	cache := make([][]int, rows+1)
	ones := make([][]int, 0)
	for r := 0; r <= rows; r++ {
		cache[r] = make([]int, cols+1)
	}

	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			one := 0
			if matrix[r][c] == '1' {
				ones = append(ones, []int{r, c})
				one = 1
			}
			cache[r][c] = cache[r][c+1] + cache[r+1][c] - cache[r+1][c+1] + one
		}
	}
	if len(ones) == 0 {
		return 0
	}

	ans := 1
	for _, pos := range ones {
		x, y := pos[0], pos[1]
		canSelect := rows - x
		if r := cols - y; r < canSelect {
			canSelect = r
		}
		for add := 1; add < canSelect; add++ {
			width := add + 1
			right := y + width
			down := x + width
			one := cache[x][y] - cache[down][y] - cache[x][right] + cache[down][right]
			if one == width*width && one > ans {
				ans = one
			}
		}
	}
	return ans
}

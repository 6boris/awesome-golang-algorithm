package Solution

func Solution(matrix [][]int, k int) int {
	rows, cols := len(matrix), len(matrix[0])
	cache := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		cache[i] = make([]int, cols+1)
	}
	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			cache[r][c] = matrix[r][c] + cache[r+1][c] + cache[r][c+1] - cache[r+1][c+1]
		}
	}
	first := true
	ans := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for re := r; re < rows; re++ {
				for rc := c; rc < cols; rc++ {
					x := cache[r][c] - cache[re+1][c] - cache[r][rc+1] + cache[re+1][rc+1]
					if x <= k {
						if first || x > ans {
							ans = x
							first = false
						}
					}
				}

			}
		}
	}
	return ans
}

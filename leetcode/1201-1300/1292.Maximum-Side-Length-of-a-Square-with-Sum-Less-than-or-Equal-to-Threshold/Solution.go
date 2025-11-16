package Solution

func Solution(mat [][]int, threshold int) int {
	rows, cols := len(mat), len(mat[0])
	cache := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		cache[i] = make([]int, cols+1)
	}

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			cache[i][j] = mat[i][j] + cache[i+1][j] + cache[i][j+1] - cache[i+1][j+1]
		}
	}

	var ret, ml int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ml = min(rows-i, cols-j)
			for k := 1; k <= ml; k++ {
				sum := cache[i][j] - cache[i+k][j] - cache[i][j+k] + cache[i+k][j+k]
				if sum <= threshold {
					ret = max(ret, k)
				}
			}
		}
	}
	return ret
}

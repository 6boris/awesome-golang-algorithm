package Solution

func Solution(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	cols := make([]int, n)
	rows := make([]int, m)
	for col := 0; col < n; col++ {
		rows[0] = mat[0][col]
		for row := 1; row < m; row++ {
			mat[row][col] += mat[row-1][col]
			rows[row] = mat[row][col]
		}

		mat[0][col] = mat[m-1][col]
		if k < m {
			mat[0][col] = mat[k][col]
		}
		for row := 1; row < m; row++ {
			end, start := row+k, row-k-1
			if end >= m {
				end = m - 1
			}

			mat[row][col] = rows[end]
			if start >= 0 {
				mat[row][col] -= rows[start]
			}
		}
	}

	for row := 0; row < m; row++ {
		cols[0] = mat[row][0]
		for col := 1; col < n; col++ {
			mat[row][col] += mat[row][col-1]
			cols[col] = mat[row][col]
		}

		mat[row][0] = mat[row][n-1]
		if k < n {
			mat[row][0] = mat[row][k]
		}
		for col := 1; col < n; col++ {
			left, right := col-k-1, col+k
			if right >= n {
				right = n - 1
			}
			mat[row][col] = cols[right]
			if left >= 0 {
				mat[row][col] -= cols[left]
			}
		}
	}

	return mat
}

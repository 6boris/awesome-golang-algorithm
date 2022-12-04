package Solution

func Solution(mat [][]int) int {
	ans := 0
	m, n := len(mat), len(mat[0])
	rowSum := make([]int, m)
	colSum := make([]int, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			rowSum[row] += mat[row][col]
			colSum[col] += mat[row][col]
		}
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if mat[row][col] == 1 && rowSum[row] == 1 && colSum[col] == 1 {
				ans++
			}
		}
	}
	return ans
}

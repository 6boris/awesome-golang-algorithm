package Solution

func Solution(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
	for i := 0; i < n; i++ {
		reversed := make([]int, n)
		for j := 0; j < n; j++ {
			reversed[j] = matrix[i][n-j-1]
		}
		matrix[i] = reversed
	}
}

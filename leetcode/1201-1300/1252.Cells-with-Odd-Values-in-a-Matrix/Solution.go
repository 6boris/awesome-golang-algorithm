package Solution

func Solution(n int, m int, indices [][]int) int {
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, m)
	}
	for _, val := range indices {
		r, c := val[0], val[1]
		for i := 0; i < m; i++ {
			mat[r][i]++
		}
		for i := 0; i < n; i++ {
			mat[i][c]++
		}
	}
	c := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j]%2 == 1 {
				c++
			}
		}
	}
	return c
}

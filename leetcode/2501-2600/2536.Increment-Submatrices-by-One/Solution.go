package Solution

func Solution(n int, queries [][]int) [][]int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}
	for _, q := range queries {
		for i := q[0]; i <= q[2]; i++ {
			for j := q[1]; j <= q[3]; j++ {
				mat[i][j]++
			}
		}
	}
	return mat
}

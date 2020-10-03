package Solution

func Solution(M [][]int) [][]int {
	m, n := len(M), len(M[0])
	N := make([][]int, m)
	for i := range N {
		N[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			N[i][j] = getSurrounding(M, i, j, m, n)
		}
	}
	return N
}

func getSurrounding(M [][]int, i, j, m, n int) int {
	total, c := M[i][j], 1
	if i-1 >= 0 {
		total += M[i-1][j]
		c++
		if j-1 >= 0 {
			total += M[i-1][j-1]
			c++
		}
		if j+1 < n {
			total += M[i-1][j+1]
			c++
		}
	}
	if i+1 < m {
		total += M[i+1][j]
		c++
		if j-1 >= 0 {
			total += M[i+1][j-1]
			c++
		}
		if j+1 < n {
			total += M[i+1][j+1]
			c++
		}
	}
	if j-1 >= 0 {
		total += M[i][j-1]
		c++
	}
	if j+1 < n {
		total += M[i][j+1]
		c++
	}
	return total / c
}

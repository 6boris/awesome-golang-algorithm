package Solution

func Solution(grid [][]int) bool {
	sum := 0
	m, n := len(grid), len(grid[0])

	row, col := make([]int, m), make([]int, n)
	tmp := 0
	for i := range m {
		tmp = 0
		for j := range n {
			tmp += grid[i][j]
			col[j] += grid[i][j]
		}
		sum += tmp
		row[i] = tmp
	}

	prefix := 0
	for i := 0; i < m-1; i++ {
		prefix += row[i]
		if sum-prefix == prefix {
			return true
		}
	}
	prefix = 0
	for i := 0; i < n-1; i++ {
		prefix += col[i]
		if sum-prefix == prefix {
			return true
		}
	}
	return false
}

package Solution

func Solution(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	if k == m*n {
		return grid
	}
	if k%n == 0 {
		for k > 0 {
			last := grid[m-1]
			for i := m - 1; i > 0; i-- {
				grid[i] = grid[i-1]
			}
			grid[0] = last
			k -= n
		}
		return grid
	}
	tmp := make([]int, 0)
	for _, val := range grid {
		tmp = append(tmp, val...)
	}
	for k > 0 {
		last := tmp[m*n-1]
		for i := m*n - 1; i > 0; i-- {
			tmp[i] = tmp[i-1]
		}
		tmp[0] = last
		k--
	}
	l := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = tmp[l]
			l++
		}
	}
	return grid
}

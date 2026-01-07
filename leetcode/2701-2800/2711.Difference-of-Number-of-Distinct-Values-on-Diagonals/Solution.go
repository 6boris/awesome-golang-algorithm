package Solution

func Solution(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ret := make([][]int, m)
	for i := range m {
		ret[i] = make([]int, n)
	}

	var (
		left, right int

		exist map[int]struct{}
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			left, right = 0, 0
			exist = map[int]struct{}{}
			for ii, jj := i-1, j-1; ii >= 0 && jj >= 0; ii, jj = ii-1, jj-1 {
				exist[grid[ii][jj]] = struct{}{}
			}
			left = len(exist)
			exist = map[int]struct{}{}
			for ii, jj := i+1, j+1; ii < m && jj < n; ii, jj = ii+1, jj+1 {
				exist[grid[ii][jj]] = struct{}{}
			}
			right = len(exist)
			ret[i][j] = left - right
			if ret[i][j] < 0 {
				ret[i][j] = -ret[i][j]
			}
		}
	}
	return ret
}

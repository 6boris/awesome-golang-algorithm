package Solution

import "math"

func minimumSum2(grid [][]int, u, d, l, r int) int {
	min_i, max_i := len(grid), 0
	min_j, max_j := len(grid[0]), 0
	for i := u; i <= d; i++ {
		for j := l; j <= r; j++ {
			if grid[i][j] == 1 {
				min_i = min(min_i, i)
				min_j = min(min_j, j)
				max_i = max(max_i, i)
				max_j = max(max_j, j)
			}
		}
	}
	if min_i <= max_i {
		return (max_i - min_i + 1) * (max_j - min_j + 1)
	}
	return math.MaxInt32 / 3
}

func rotate(vec [][]int) [][]int {
	n, m := len(vec), len(vec[0])
	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ret[m-j-1][i] = vec[i][j]
		}
	}
	return ret
}

func solve(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	res := n * m
	for i := 0; i+1 < n; i++ {
		for j := 0; j+1 < m; j++ {
			res = min(res, minimumSum2(grid, 0, i, 0, m-1)+
				minimumSum2(grid, i+1, n-1, 0, j)+
				minimumSum2(grid, i+1, n-1, j+1, m-1))
			res = min(res, minimumSum2(grid, 0, i, 0, j)+
				minimumSum2(grid, 0, i, j+1, m-1)+
				minimumSum2(grid, i+1, n-1, 0, m-1))
		}
	}
	for i := 0; i+2 < n; i++ {
		for j := i + 1; j+1 < n; j++ {
			res = min(res, minimumSum2(grid, 0, i, 0, m-1)+
				minimumSum2(grid, i+1, j, 0, m-1)+
				minimumSum2(grid, j+1, n-1, 0, m-1))
		}
	}
	return res
}

func Solution(grid [][]int) int {
	rgrid := rotate(grid)
	return min(solve(grid), solve(rgrid))
}

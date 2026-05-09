package Solution

import "sort"

func Solution(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ret := make([][]int, m-k+1)
	for i := range ret {
		ret[i] = make([]int, n-k+1)
	}

	cache := make([]int, k*k)
	ii := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i+k > m || j+k > n {
				continue
			}
			if k == 1 {
				ret[i][j] = 0
				continue
			}
			ii = 0
			for x := i; x < i+k; x++ {
				for y := j; y < j+k; y++ {
					cache[ii] = grid[x][y]
					ii++
				}
			}
			sort.Ints(cache)
			cmp := 0x7fffffff
			for idx := 1; idx < ii; idx++ {
				if cache[idx] == cache[idx-1] {
					continue
				}
				cmp = min(cmp, cache[idx]-cache[idx-1])
			}
			if cmp == 0x7fffffff {
				cmp = 0
			}
			ret[i][j] = cmp
		}
	}
	return ret
}

package Solution

import "maps"

func Solution(grid [][]int) bool {
	sum := 0
	m, n := len(grid), len(grid[0])

	row, col := make([]int, m), make([]int, n)
	tmp := 0
	cache := map[int]int{}

	for i := range m {
		tmp = 0
		for j := range n {
			cache[grid[i][j]]++
			tmp += grid[i][j]
			col[j] += grid[i][j]
		}
		sum += tmp
		row[i] = tmp
	}

	used := map[int]int{}
	cmp := maps.Clone(cache)
	prefix := 0
	for i := 0; i < m-1; i++ {
		prefix += row[i]
		left := sum - prefix
		if left == prefix {
			return true
		}
		for j := 0; j < n; j++ {
			used[grid[i][j]]++
			cmp[grid[i][j]]--
		}

		diff := prefix - left
		if diff < 0 {
			if cmp[-diff] > 0 {
				if n == 1 {
					if grid[m-1][0] == -diff || grid[i+1][0] == -diff {
						return true
					}
				} else if i < m-2 || (grid[m-1][0] == -diff || grid[m-1][n-1] == -diff) {
					return true
				}
			}
		} else if used[diff] > 0 {
			if n == 1 {
				if grid[0][0] == diff || grid[i][0] == diff {
					return true
				}
			} else if i >= 1 || grid[0][0] == diff || grid[0][n-1] == diff {
				return true
			}
		}

	}
	cmp = maps.Clone(cache)
	used = map[int]int{}
	prefix = 0
	for i := 0; i < n-1; i++ {
		prefix += col[i]
		left := sum - prefix
		if left == prefix {
			return true
		}
		for j := 0; j < m; j++ {
			used[grid[j][i]]++
			cmp[grid[j][i]]--
		}
		diff := prefix - left
		if diff < 0 {
			if cmp[-diff] > 0 {
				if m == 1 {
					if grid[0][n-1] == -diff || grid[0][i+1] == -diff {
						return true
					}
				} else if i < n-2 || (grid[0][n-1] == -diff || grid[m-1][n-1] == -diff) {
					return true
				}
			}
		} else if used[diff] > 0 {
			if m == 1 {
				if grid[0][0] == diff || grid[0][i] == diff {
					return true
				}
			} else if i >= 1 || (grid[0][0] == diff || grid[m-1][0] == diff) {
				return true
			}
		}
	}
	return false
}

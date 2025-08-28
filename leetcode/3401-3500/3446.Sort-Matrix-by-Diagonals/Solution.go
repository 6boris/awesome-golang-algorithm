package Solution

import "sort"

func Solution(grid [][]int) [][]int {
	n := len(grid)
	for i := 0; i < n; i++ {
		// j = 0
		have := []int{}
		for r, j := i, 0; r < n; r, j = r+1, j+1 {
			have = append(have, grid[r][j])
		}
		sort.Slice(have, func(i, j int) bool {
			return have[i] > have[j]
		})
		index := 0
		for r, j := i, 0; r < n; r, j, index = r+1, j+1, index+1 {
			grid[r][j] = have[index]
		}
	}
	for j := 1; j < n; j++ {
		// i = 1
		have := []int{}
		for i, r := 0, j; r < n; i, r = i+1, r+1 {
			have = append(have, grid[i][r])
		}
		sort.Ints(have)
		for i, r, index := 0, j, 0; r < n; i, r, index = i+1, r+1, index+1 {
			grid[i][r] = have[index]
		}
	}
	return grid
}

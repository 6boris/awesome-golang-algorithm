package Solution

import "sort"

func Solution(grid [][]int, x int) int {
	list := []int{}
	rows, cols := len(grid), len(grid[0])
	for i := range rows {
		for j := range cols {
			list = append(list, grid[i][j])
		}
	}
	sort.Ints(list)
	mid := list[len(list)/2]
	ans := 0
	for _, n := range list {
		if n%x != mid%x {
			return -1
		}
		diff := n - mid
		if diff < 0 {
			diff = -diff
		}
		ans += diff / x
	}
	return ans
}

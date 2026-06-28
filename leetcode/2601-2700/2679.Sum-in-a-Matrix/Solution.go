package Solution

import "sort"

func Solution(nums [][]int) int {
	n := len(nums[0])
	var ret int
	col := make([]int, n)
	for i := range nums {
		sort.Ints(nums[i])
		for j := range n {
			col[j] = max(col[j], nums[i][j])
		}
	}
	for i := range col {
		ret += col[i]
	}
	return ret
}

package Solution

import "sort"

func Solution(cost []int) int {
	sort.Ints(cost)
	sum, n := 0, len(cost)
	for i := range n {
		sum += cost[i]
	}
	// 0, 1, 2, 3, 4, 5
	for i := n - 3; i >= 0; i -= 3 {
		sum -= cost[i]
	}
	return sum
}

package Solution

import "sort"

func Solution(piles []int) int {
	// 1, 2, 2, 4, 7, 8
	sort.Ints(piles)
	l, r := 0, len(piles)-2
	ans := 0
	for ; l < r; l, r = l+1, r-2 {
		ans += piles[r]
	}
	return ans
}

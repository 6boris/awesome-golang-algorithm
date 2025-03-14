package Solution

import (
	"slices"
	"sort"
)

func Solution(candies []int, k int64) int {
	m := slices.Max(candies)
	var ok func(int) bool
	ok = func(n int) bool {
		cnt := int64(0)
		for _, c := range candies {
			cnt += int64(c / n)
		}
		return cnt >= k
	}
	// 0, 1, 2, 3, ,4
	// 1, 2, 3, 4, ,5
	x := sort.Search(m+1, func(i int) bool {
		return !ok(i + 1)
	})
	return x
}

package Solution

import (
	"slices"
	"sort"
)

func Solution(nums []int, threshold int) int {
	m := slices.Max(nums)
	var ok func(int) int
	ok = func(divisor int) int {
		sum := 0
		for _, n := range nums {
			a := n / divisor
			if n%divisor != 0 {
				a++
			}
			sum += a
		}
		return sum
	}
	idx := sort.Search(m+1, func(n int) bool {
		if n == 0 {
			return false
		}
		return ok(n) <= threshold
	})
	return idx
}

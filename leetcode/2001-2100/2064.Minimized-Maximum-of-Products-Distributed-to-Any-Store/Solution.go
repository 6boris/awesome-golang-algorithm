package Solution

import "slices"

func Solution(n int, quantities []int) int {
	m := slices.Max(quantities)
	if n == len(quantities) {
		return m
	}
	var ok func(int) bool
	ok = func(x int) bool {
		need := n
		for _, n := range quantities {
			a := n / x
			if n%x != 0 {
				a++
			}
			if need < a {
				return false
			}
			need -= a
		}
		return true
	}
	left, right := 1, m+1
	ans := 0
	for left < right {
		mid := (left + right) / 2
		if ok(mid) {
			right = mid
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}

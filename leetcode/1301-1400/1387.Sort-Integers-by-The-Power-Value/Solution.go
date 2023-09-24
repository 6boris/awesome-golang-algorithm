package Solution

import (
	"sort"
)

func power(x int, cache map[int]int) int {
	if x == 1 {
		return 1
	}
	if v, ok := cache[x]; ok {
		return v
	}
	var ans int
	if x&1 == 0 {
		ans = power(x/2, cache) + 1
	} else {
		ans = power(x*3+1, cache) + 1
	}
	cache[x] = ans
	return ans
}
func Solution(lo, hi, k int) int {
	cache := make(map[int]int)
	x := make([][2]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		x[i-lo] = [2]int{power(i, cache), i}
	}
	sort.SliceStable(x, func(i, j int) bool {
		return x[i][0] < x[j][0]
	})
	return x[k-1][1]
}

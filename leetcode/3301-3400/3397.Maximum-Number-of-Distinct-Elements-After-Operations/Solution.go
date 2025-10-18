package Solution

import (
	"math"
	"sort"
)

func Solution(nums []int, k int) int {
	sort.Ints(nums)
	cnt := 0
	prev := math.MinInt32

	for _, num := range nums {
		curr := min(max(num-k, prev+1), num+k)
		if curr > prev {
			cnt++
			prev = curr
		}
	}
	return cnt
}

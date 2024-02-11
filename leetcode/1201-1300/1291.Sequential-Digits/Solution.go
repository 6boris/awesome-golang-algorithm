package Solution

import "sort"

func Solution(low int, high int) []int {
	start := []int{12, 23, 34, 45, 56, 67, 78, 89}
	ans := make([]int, 0)
	for i, n := range start {
		now, l := n, i+2
		for now < low && l < 10 {
			l++
			now = now*10 + l
		}
		for now <= high && l < 10 {
			ans = append(ans, now)
			l++
			now = now*10 + l
		}
	}
	sort.Ints(ans)
	return ans
}

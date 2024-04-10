package Solution

import "sort"

func Solution(deck []int) []int {
	// 长度是奇数还是偶数
	sort.Ints(deck)
	skip := false
	ans := make([]int, len(deck))
	a, b := 0, 0
	for a < len(deck) {
		if ans[b] == 0 {
			if !skip {
				ans[b] = deck[a]
				a++
			}
			skip = !skip
		}
		b = (b + 1) % len(deck)
	}
	return ans
}

package Solution

import "sort"

func Solution(chalk []int, k int) int {
	l := len(chalk)
	for i := 1; i < l; i++ {
		chalk[i] += chalk[i-1]
	}
	k %= chalk[l-1]
	idx := sort.Search(l, func(i int) bool {
		return chalk[i] >= k
	})
	if idx == l {
		return 0
	}
	if chalk[idx] == k {
		return (idx + 1) % l
	}
	return idx
}

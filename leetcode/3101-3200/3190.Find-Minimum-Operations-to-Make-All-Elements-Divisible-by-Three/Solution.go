package Solution

import "sort"

func Solution(nums []int) int {

	byThree := []int{0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51}
	l := len(byThree)
	var ret int

	for _, n := range nums {
		index := sort.Search(l, func(i int) bool {
			return byThree[i] >= n
		})
		ret += min(byThree[index]-n, n-byThree[index-1])
	}

	return ret
}

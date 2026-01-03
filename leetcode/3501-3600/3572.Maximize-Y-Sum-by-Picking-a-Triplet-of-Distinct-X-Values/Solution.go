package Solution

import (
	"sort"
)

func Solution(x []int, y []int) int {
	maxItem := make(map[int]int)
	for index, n := range x {
		maxItem[n] = max(maxItem[n], y[index])
	}
	if len(maxItem) < 3 {
		return -1
	}
	values := make([]int, 0)
	for _, v := range maxItem {
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	ret := 0
	for i := range 3 {
		ret += values[i]
	}
	return ret
}

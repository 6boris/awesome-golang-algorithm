package Solution

import "slices"

func Solution(values []int) int {
	// 8, 1, 5, 2, 6
	// 0, 7, 6, 5, 4
	// 3, 4, 4, 5, 0
	l := len(values)
	left, right := make([]int, l), make([]int, l)
	for i := 1; i < l; i++ {
		left[i] = max(left[i-1]-1, values[i-1]-1)
	}
	for i := l - 2; i >= 0; i-- {
		right[i] = max(right[i+1]-1, values[i+1]-1)
	}
	for i := 0; i < l; i++ {
		left[i] = max(left[i], right[i]) + values[i]
	}
	return slices.Max(left)
}

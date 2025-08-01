package Solution

import "sort"

func Solution(num int) bool {
	l := (1 << 31) - 1
	index := sort.Search(l, func(i int) bool {
		return i*i >= num
	})
	return index*index == num
}

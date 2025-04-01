package Solution

import "sort"

func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

func mySqrt1(x int) int {
	idx := sort.Search(x, func(i int) bool {
		return i*i >= x
	})
	if idx*idx == x {
		return idx
	}
	return idx - 1
}

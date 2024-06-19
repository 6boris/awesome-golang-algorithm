package Solution

import "sort"

func Solution(bloomDay []int, m int, k int) int {
	l := len(bloomDay)
	need := m * k
	if need > l {
		return -1
	}
	var ok func(int) bool
	ok = func(day int) bool {
		c := 0
		complete := 0
		for i := 0; i < l; i++ {
			if bloomDay[i] <= day {
				c++
				continue
			}
			complete += c / k
			c = 0
		}
		complete += c / k
		return complete >= m
	}

	dst := make([]int, 0)
	in := make(map[int]struct{})
	for _, n := range bloomDay {
		if _, ok := in[n]; !ok {
			dst = append(dst, n)
			in[n] = struct{}{}
		}
	}
	sort.Ints(dst)
	left, right := 0, len(dst)
	ans := -1
	for left < right {
		mid := (right + left) / 2
		if ok(dst[mid]) {
			ans = dst[mid]
			right = mid
			continue
		}
		left = mid + 1
	}
	return ans
}

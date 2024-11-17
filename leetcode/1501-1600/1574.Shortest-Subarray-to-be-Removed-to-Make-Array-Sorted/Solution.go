package Solution

import (
	"sort"
)

func Solution(arr []int) int {
	if sort.IsSorted(sort.IntSlice(arr)) {
		return 0
	}
	l := len(arr)
	ll, rr := make([]int, l), make([]int, l)
	ll[0] = 1
	for i := 1; i < l; i++ {
		if arr[i] >= arr[i-1] {
			ll[i] = ll[i-1] + 1
			continue
		}
		ll[i] = 1
	}
	rr[l-1] = 1
	for i := l - 2; i >= 0; i-- {
		if arr[i] <= arr[i+1] {
			rr[i] = rr[i+1] + 1
			continue
		}
		rr[i] = 1
	}

	var ok func(int) bool
	ok = func(length int) bool {
		start, end := 0, length-1
		for ; end < l; start, end = start+1, end+1 {
			lok := (start == 0 || ll[start-1] == start)
			rok := (end == l-1 || rr[end+1] == l-end-1)
			if lok && rok {
				if start == 0 || end == l-1 || (arr[start-1] <= arr[end+1]) {
					return true
				}
				continue
			}
		}
		return false
	}
	left, right := 1, l+1
	ans := l - 1
	for left < right {
		mid := (left + right) / 2
		if ok(mid) {
			right = mid
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}

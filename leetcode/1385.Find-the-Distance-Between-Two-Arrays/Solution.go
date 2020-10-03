package Solution

import "sort"

func Solution(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	c := 0
	for _, v1 := range arr1 {
		l, r := 0, len(arr2)-1
		for l <= r {
			mid := (l + r) / 2
			if abs(v1-arr2[mid]) <= d {
				c++
				break
			}
			if arr2[mid] > v1 {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return len(arr1) - c
}

func abs(val int) int {
	if val < 0 {
		return -val
	} else {
		return val
	}
}

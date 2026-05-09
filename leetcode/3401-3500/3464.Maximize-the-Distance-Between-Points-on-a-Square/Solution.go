package Solution

import "sort"

func Solution(side int, points [][]int, k int) int {
	arr := make([]int64, 0, len(points))

	for _, p := range points {
		x, y := p[0], p[1]
		if x == 0 {
			arr = append(arr, int64(y))
		} else if y == side {
			arr = append(arr, int64(side+x))
		} else if x == side {
			arr = append(arr, int64(side*3-y))
		} else {
			arr = append(arr, int64(side*4-x))
		}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	lo, hi := int64(1), int64(side)
	ans := 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if check(arr, int64(side), k, mid) {
			lo = mid + 1
			ans = int(mid)
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func check(arr []int64, side int64, k int, limit int64) bool {
	perimeter := side * 4

	for _, start := range arr {
		end := start + perimeter - limit
		cur := start

		for i := 0; i < k-1; i++ {
			idx := lowerBound(arr, cur+limit)
			if idx == len(arr) || arr[idx] > end {
				cur = -1
				break
			}
			cur = arr[idx]
		}

		if cur >= 0 {
			return true
		}
	}
	return false
}

func lowerBound(arr []int64, target int64) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

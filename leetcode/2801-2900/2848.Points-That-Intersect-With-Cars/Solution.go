package Solution

import "sort"

func Solution(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})

	points := 0
	var (
		left, right int
		first       = true
	)
	for _, car := range nums {
		if first {
			left, right = car[0], car[1]
			points += right - left + 1
			first = false
			continue
		}
		if car[0] >= right {
			points += car[1] - car[0]
			if car[0] > right {
				points++
			}
			left, right = car[0], car[1]
			continue
		}
		if car[1] > right {
			points += car[1] - right
			right = car[1]
		}
	}
	return points
}

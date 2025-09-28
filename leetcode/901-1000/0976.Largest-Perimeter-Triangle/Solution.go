package Solution

import "sort"

func Solution(nums []int) int {
	sort.Ints(nums)
	// 1, 1, 2, 10
	l := len(nums)
	for i := l - 1; i >= 2; i-- {
		a, b, c := nums[i], nums[i-1], nums[i-2]
		if a+b > c && a+c > b && b+c > a {
			return a + b + c
		}
	}
	return 0
}

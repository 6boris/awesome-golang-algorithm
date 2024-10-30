package Solution

import "sort"

func isPerfectSquare(num int) int {
	for i := 1; i*i <= num; i++ {
		if i*i == num {
			return i
		}
	}
	return -1
}

func Solution(nums []int) int {
	sort.Ints(nums)
	ans := -1
	pre := make(map[int]int)
	for _, n := range nums {
		pre[n] = 1
	}
	for _, n := range nums {
		i := isPerfectSquare(n)
		if i == -1 {
			continue
		}
		if l, ok := pre[i]; ok {
			pre[n] = l + 1
			ans = max(ans, pre[n])
		}
	}
	return ans
}

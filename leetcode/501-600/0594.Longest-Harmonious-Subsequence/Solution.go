package Solution

import "sort"

func Solution(nums []int) int {
	keys := make(map[int]int)
	a := make([]int, 0)
	for _, n := range nums {
		if _, ok := keys[n]; !ok {
			a = append(a, n)
		}
		keys[n]++
	}
	sort.Ints(a)
	ans := 0
	for i := 1; i < len(a); i++ {
		if a[i]-a[i-1] == 1 && keys[a[i]]+keys[a[i-1]] > ans {
			ans = keys[a[i]] + keys[a[i-1]]
		}
	}
	return ans
}

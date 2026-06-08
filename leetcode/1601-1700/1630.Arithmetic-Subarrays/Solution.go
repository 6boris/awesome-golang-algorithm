package Solution

import "sort"

func Solution(nums, l, r []int) []bool {
	ans := make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		s, e := l[i], r[i]
		c := make([]int, e-s+1)
		copy(c, nums[s:e+1])
		sort.Ints(c)
		ok := true
		diff := -1
		for i := 1; i < len(c); i++ {
			if diff != -1 && c[i]-c[i-1] != diff {
				ok = false
				break
			}
			diff = c[i] - c[i-1]
		}
		ans[i] = ok
	}
	return ans
}

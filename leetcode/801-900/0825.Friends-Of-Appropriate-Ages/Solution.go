package Solution

import "sort"

func Solution(ages []int) int {
	count := make(map[int]int)
	distinctAges := make([]int, 0)
	for _, n := range ages {
		count[n]++
		if count[n] == 1 {
			distinctAges = append(distinctAges, n)
		}
	}
	sort.Ints(distinctAges)

	ans := 0
	for i := 1; i < len(distinctAges); i++ {
		c := 0
		for pre := i - 1; pre >= 0; pre-- {
			if distinctAges[pre] <= distinctAges[i]/2+7 {
				continue
			}
			if distinctAges[pre] > 100 && distinctAges[i] < 100 {
				continue
			}
			c += count[distinctAges[pre]]
		}
		ans += c * count[distinctAges[i]]
	}
	for n, c := range count {
		if c == 1 {
			continue
		}
		if n <= n/2+7 {
			continue
		}
		ans += (c - 1) * c
	}

	return ans
}

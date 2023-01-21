package Solution

import "sort"

func Solution(capacity []int, rocks []int, additionalRocks int) int {
	diff := make([]int, len(capacity))
	for i := 0; i < len(capacity); i++ {
		diff[i] = capacity[i] - rocks[i]
	}
	sort.Ints(diff)
	i := 0
	ans := 0
	for ; i < len(capacity) && additionalRocks > 0; i++ {
		if diff[i] <= 0 {
			ans++
			continue
		}
		if diff[i] <= additionalRocks {
			ans++
			additionalRocks -= diff[i]
		}
	}
	return ans
}

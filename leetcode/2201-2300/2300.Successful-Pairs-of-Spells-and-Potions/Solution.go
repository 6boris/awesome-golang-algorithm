package Solution

import "sort"

func Solution(spells []int, potions []int, success int64) []int {
	ans := make([]int, len(spells))
	sort.Ints(potions)
	for i, n := range spells {
		idx := 0
		for ; idx < len(potions); idx++ {
			if int64(potions[idx])*int64(n) >= success {
				break
			}
		}
		ans[i] = len(potions) - idx
	}
	return ans
}

package Solution

import "sort"

func Solution(amount int, coins []int) int {
	result := make([]int, amount+1)
	tmp := make([]int, amount+1)
	sort.Ints(coins)
	result[0] = 1
	baseCoin := coins[0]
	for idx := baseCoin; idx <= amount; idx += baseCoin {
		result[idx] = 1
	}
	for idx := 1; idx < len(coins); idx++ {
		for start := 0; start <= amount; start++ {
			for walker := 0; walker+start <= amount; walker += coins[idx] {
				tmp[walker+start] += result[start]
			}
		}
		for idx := 0; idx <= amount; idx++ {
			result[idx], tmp[idx] = tmp[idx], 0
		}
	}
	return result[amount]
}

package Solution

import "math"

/**
持有一股
买卖一次
*/

/**
Brute Force
时间复杂度O(n^2)
空间复杂度O(1)
*/
func maxProfit1(prices []int) int {
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxprofit {
				maxprofit = profit
			}
		}
	}
	return maxprofit
}

//	Greedy
//	扫描一遍
func maxProfit2(prices []int) int {
	min := math.MaxInt32
	max := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if max < prices[i]-min {
			max = prices[i] - min
		}
	}

	return max
}

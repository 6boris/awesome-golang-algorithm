package Solution

import "math"

func maxProfit(prices []int) int {
	minCost, maxBenefit := math.MaxInt32, 0
	for _, price := range prices {
		minCost = min(minCost, price)
		maxBenefit = max(maxBenefit, price-minCost)
	}
	return maxBenefit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

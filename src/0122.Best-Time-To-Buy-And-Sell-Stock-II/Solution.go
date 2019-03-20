package Solution

import (
	"math"
)

/**
持有一股
买卖无数次
*/

//	Greedy
func maxProfit(prices []int) int {
	maxprofix := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxprofix += prices[i] - prices[i-1]
		}
	}
	return maxprofix
}

//	DP
//dp[i][0] means the maximum profit after sunset of day i with no stock in hand
//dp[i][1] means the maximum profit after sunset of day i with stock in hand
//dp[i][0] = max(dp[i-1][1]+price[i],dp[i-1][0]) --> Sell on day i or do nothing
//dp[i][1] = max(dp[i-1][0])-price[i],dp[i-1][1]) --> Buy on day i or do nothing
//dp[0][0]=0,dp[0][1]=INT_MIN
func maxProfit2(prices []int) int {
	// 初始化DP
	n := len(prices)
	dp := make([][]int, 0)
	for i := 0; i <= n; i++ {
		dp = append(dp, make([]int, 2))
	}
	dp[0][0], dp[0][1] = 0, math.MinInt32
	profit_max := 0
	for i := 0; i < n; i++ {
		dp[i+1][0] = max(dp[i][1]+prices[i], dp[i][0])
		dp[i+1][1] = max(dp[i][0]-prices[i], dp[i][1])

		profit_max = max(profit_max, dp[i+1][0])
		profit_max = max(profit_max, dp[i+1][1])
	}
	return profit_max
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//	DFS

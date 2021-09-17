package Solution

/**
持有一股
买卖无数次
*/

// Greedy
func maxProfit_1(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// DP
// dp[i][0] means the maximum profit after sunset of day i with no stock in hand
// dp[i][1] means the maximum profit after sunset of day i with stock in hand
// dp[i][0] = max(dp[i-1][1]+price[i],dp[i-1][0]) --> Sell on day i or do nothing
// dp[i][1] = max(dp[i-1][0]-price[i],dp[i-1][1]) --> Buy on day i or do nothing
// dp[0][0]=0,dp[0][1]=INT_MIN

func maxProfit_2(prices []int) int {
	// 初始化DP
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

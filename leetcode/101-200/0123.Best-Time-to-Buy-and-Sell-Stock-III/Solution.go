package Solution

/*
操作
	dp[i][0] 没有操作

	dp[i][1] 第一次买入
		以前已经买入｜以前未买入、今天买入
	dp[i][2] 第一次卖出
		以前已经卖出｜以前未卖出、今天卖出
	dp[i][3] 第二次买入
		...
	dp[i][4] 第二次卖出
		...
*/
func maxProfit_1(prices []int) int {
	dp, n := make([][5]int, len(prices)), len(prices)
	dp[0][1], dp[0][3] = -prices[0], -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[n-1][4]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

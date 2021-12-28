package Solution

/*
状态定义
	有股票 - 0
	无股票
		以前卖出、无冷冻 - 1
		今天卖出、有冷冻 - 2
		以前卖出、冷冻中 - 3

状态转移
	今天		昨天
	0
		<- 0
		<- 1
		<- 3
	1
		<- 1
		<- 3
	2
		<- 0
	3
		<- 2
*/
func maxProfit_1(prices []int) int {
	dp, n := make([][4]int, len(prices)), len(prices)
	dp[0][0], dp[0][1] = -prices[0], 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3], dp[i-1][1])-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[n-1][3], max(dp[n-1][1], dp[n-1][2]))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package Solution

func maxProfit_1(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

/*
当前有无股票
    有
        昨天有、今天不动
        昨天无、今天买
    无
        昨天有、今天卖
        昨天无、今天不动
*/
func maxProfit_2(prices []int) int {
	dp, n := make([][2]int, len(prices)), len(prices)
	dp[0][0], dp[0][1] = -prices[0], 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[n-1][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

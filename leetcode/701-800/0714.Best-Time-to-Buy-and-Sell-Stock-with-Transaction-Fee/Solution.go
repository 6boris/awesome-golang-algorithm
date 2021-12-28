package Solution

func maxProfit_1(prices []int, fee int) int {
	ans, minPrice := 0, prices[0]
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if v > minPrice && v > minPrice+fee {
			ans += v - minPrice - fee
			minPrice = v - fee
		}
	}
	return ans
}

func maxProfit_2(prices []int, fee int) int {
	dp, n := make([][2]int, len(prices)), len(prices)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

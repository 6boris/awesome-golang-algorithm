package Solution

/*
[天数][股票状态]
股票状态: 奇数表示第 k 次交易持有/买入, 偶数表示第 k 次交易不持有/卖出, 0 表示没有操作
*/

func maxProfit_1(k int, prices []int) int {
	n, dp := len(prices), [][]int{}
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, k*2+1))
	}
	for i := 1; i < k*2; i += 2 {
		dp[0][i] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < k*2; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[n-1][k*2]
}

func maxProfit_2(k int, prices []int) int {
	n := len(prices)
	dp := make([]int, k*2+1)
	for i := 1; i < k*2; i += 2 {
		dp[i] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < k*2; j += 2 {
			dp[j+1] = max(dp[j+1], dp[j]-prices[i])
			dp[j+2] = max(dp[j+2], dp[j+1]+prices[i])
		}
	}
	return dp[k+2]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

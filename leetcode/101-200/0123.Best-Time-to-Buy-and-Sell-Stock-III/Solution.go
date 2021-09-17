package Solution

import "math"

func maxProfit2(prices []int) int {
	t1_b, t1_s, t2_b, t2_s := math.MinInt32, 0, math.MinInt32, 0

	for _, v := range prices {
		t1_b = max(t1_b, 0-v)
		t1_s = max(t1_s, t1_b+v)
		t2_b = max(t2_b, t1_s-v)
		t2_s = max(t2_s, t2_b+v)
	}
	return t2_s
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return y
}

func maxProfit(prices []int) int {
	g, f := make([]int, len(prices)), make([]int, len(prices))
	ans, n, low := 0, len(prices), prices[0]

	for i := 1; i < n; i++ {
		low = min(low, prices[i])
		f[i] = max(f[i-1], prices[i]-low)
	}
	high := prices[n-1]
	for i := n - 2; i >= 0; i-- {
		high = max(high, prices[i])
		g[i] = max(g[i+1], high-prices[i])
	}

	for i := 0; i < n; i++ {
		ans = max(ans, f[i]+g[i])
	}
	return ans
}

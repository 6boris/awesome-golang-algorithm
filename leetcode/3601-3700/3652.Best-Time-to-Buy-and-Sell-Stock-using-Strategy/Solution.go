package Solution

func Solution(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	profitSum := make([]int64, n+1)
	priceSum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		profitSum[i+1] = profitSum[i] + int64(prices[i])*int64(strategy[i])
		priceSum[i+1] = priceSum[i] + int64(prices[i])
	}
	res := profitSum[n]
	for i := k - 1; i < n; i++ {
		leftProfit := profitSum[i-k+1]
		rightProfit := profitSum[n] - profitSum[i+1]
		changeProfit := priceSum[i+1] - priceSum[i-k/2+1]
		res = max(res, leftProfit+changeProfit+rightProfit)
	}
	return res
}

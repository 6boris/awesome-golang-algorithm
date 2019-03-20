package Solution

func maxProfit(k int, prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	// 可允许交易次数大于size，可以获得每一次上涨的利润
	if k >= size {
		return profits(prices)
	}
	// 另一个是当前到达第i天，最多可进行j次交易，并且最后一次交易在当天卖出的最好的利润是多少
	local := make([]int, k+1)
	// 当前到达第i天可以最多进行j次交易，最好的利润是多少
	global := make([]int, k+1)

	for i := 1; i < size; i++ {
		diff := prices[i] - prices[i-1]
		for j := k; j >= 1; j-- {
			local[j] = max(global[j-1]+max(diff, 0), local[j]+diff)
			global[j] = max(local[j], global[j])
		}
	}

	return global[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func profits(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			res += tmp
		}
	}

	return res
}

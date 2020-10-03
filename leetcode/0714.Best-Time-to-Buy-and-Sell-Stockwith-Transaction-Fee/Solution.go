package Solution

func maxProfit(prices []int, fee int) int {
	cash, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		cash = max(cash, hold+prices[i]-fee)
		hold = max(hold, cash-prices[i])
	}
	return cash
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

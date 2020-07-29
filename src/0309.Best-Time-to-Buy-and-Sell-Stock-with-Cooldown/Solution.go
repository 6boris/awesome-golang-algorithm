package Solution

func maxProfit(prices []int) int {
	sold, held, reset := -1<<31, -1<<31, 0

	for _, price := range prices {
		preSold := sold
		sold = held + price
		held = max(held, reset-price)
		reset = max(reset, preSold)

	}
	return max(sold, reset)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

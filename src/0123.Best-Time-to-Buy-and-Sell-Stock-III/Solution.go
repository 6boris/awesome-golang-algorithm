package Solution

import "math"

func maxProfit(prices []int) int {
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

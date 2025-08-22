package Solution

func Solution(prices []int) int64 {
	l := len(prices)
	c := int64(1)
	ans := c
	for i := 1; i < l; i++ {
		if prices[i] == prices[i-1]-1 {
			c++
		} else {
			c = 1
		}
		ans += c
	}
	return ans
}

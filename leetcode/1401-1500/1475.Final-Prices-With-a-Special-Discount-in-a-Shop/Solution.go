package Solution

func Solution(prices []int) []int {
	l := len(prices)
	dst := make([]int, l)
	copy(dst, prices)
	for i := l - 2; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			if prices[j] <= prices[i] {
				dst[i] -= prices[j]
				break
			}
		}
	}
	return dst
}

package Solution

func Solution(cardPoints []int, k int) int {
	length := len(cardPoints)
	r := 0
	for idx := length - 1; idx >= length-k; idx-- {
		r += cardPoints[idx]
	}

	ans, l := r, 0
	for idx := 0; idx < k; idx++ {
		l += cardPoints[idx]
		r -= cardPoints[length-k+idx]
		if l+r > ans {
			ans = l + r
		}
	}
	return ans
}

package Solution

func Solution(rating []int) int {
	ans, n := 0, len(rating)
	dpG, dpL := make([]int, n), make([]int, n)
	for i, r := range rating {
		for j := 0; j < i; j++ {
			if r > rating[j] {
				dpG[i]++
				ans += dpG[j]
			} else if r < rating[j] {
				dpL[i]++
				ans += dpL[j]
			}
		}
	}
	return ans
}

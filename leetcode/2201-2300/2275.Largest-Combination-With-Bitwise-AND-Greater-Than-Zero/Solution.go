package Solution

func Solution(candidates []int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		tmp := 0
		for j := range candidates {
			if candidates[j]&1 == 1 {
				tmp++
			}
			candidates[j] >>= 1
		}
		ans = max(ans, tmp)
	}
	return ans
}

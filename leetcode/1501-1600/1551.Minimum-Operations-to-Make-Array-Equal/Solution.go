package Solution

func Solution(n int) int {
	ans := 0
	for i := 0; i < n; i++ {
		r := 2*i + 1
		if r >= n {
			break
		}
		ans += n - r
	}
	return ans
}

package Solution

func lastRemaining(n int, m int) int {
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (m + ans) % i
	}
	return ans
}

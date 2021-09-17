package Solution

func nthUglyNumber(n int) int {
	a, b, c, dp := 0, 0, 0, make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(min(n2, n3), n5)
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	return dp[n-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

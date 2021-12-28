package Solution

//	优化后
func climbStairs_1(n int) int {
	a, b, c := 0, 0, 1
	for i := 1; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}

//	标准斐波拉契数列
func climbStairs_2(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

var m = make(map[int]int)

func climbStairs_3(n int) int {

	if n <= 2 {
		return n
	}
	if v, ok := m[n]; ok {
		return v
	} else {
		m[n] = climbStairs_3(n-1) + climbStairs_3(n-2)
		return m[n]
	}
}

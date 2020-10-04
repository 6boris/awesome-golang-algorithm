package Solution

//	递归解法
func jumpFloor(n int) int {
	if n <= 1 {
		return 1
	}
	return jumpFloor(n-1) + jumpFloor(n-2)
}

//	记忆化搜索
func jumpFloor2(n int) int {
	m := make([]int, 45, 45)
	return jumpFloorRecursion(n, m)
}

func jumpFloorRecursion(n int, m []int) int {
	if n <= 1 {
		return 1
	}
	if m[n] != 0 {
		return m[n]
	} else {
		m[n] = jumpFloorRecursion(n-1, m) + jumpFloorRecursion(n-2, m)
	}
	return m[n]
}

//	动态规划
func jumpFloor3(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//	动态规划 [优化空间]
func jumpFloor4(n int) int {
	if n <= 1 {
		return n
	}
	a, b, c := 1, 1, 0

	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

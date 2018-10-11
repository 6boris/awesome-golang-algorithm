package Solution

//	优化后
func climbStairs(n int) int {
	a, b := 1, 1

	for n > 0 {
		b += a
		a = b - a
		n--
	}
	return a
}

//	标准斐波拉契数列
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	aux := make([]int, n+1)
	aux[1] = 1
	aux[2] = 2
	for i := 3; i <= n; i++ {
		aux[i] = aux[i-1] + aux[i-2]
	}
	return aux[n]
}

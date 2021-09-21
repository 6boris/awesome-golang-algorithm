package Solution

func fib_1(n int) int {
	if n < 2 {
		return n
	}
	return fib_1(n-1) + fib_1(n-2)
}

var fM = map[int]int{}

func fib_2(n int) int {
	if n < 2 {
		return n
	}
	if _, ok := fM[n]; ok {
		return fM[n]
	}
	fM[n] = fib_2(n-1) + fib_2(n-2)
	return fM[n]
}

func fib_3(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib_4(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

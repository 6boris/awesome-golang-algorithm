package Solution

const mod1359 = 1000000007

func Solution(n int) int {
	// p1, d1, p2, d2
	// 1 -> p1d1 = 1
	// 2 ->
	// p2d2p1d1, p2p1d2d1, p2p1d1d2
	// p1p2d2d1, p1p2d1d2, p1d1p2d2 = 6
	// 3 - > p2 d2 p1 d1
	// 5 + 4 + 3+2+1 = 15 * 6 = 90
	sum := make([]int, 2*n)
	for i := 1; i <= 2*n-1; i++ {
		sum[i] = sum[i-1] + i
	}

	cache := make([]int, n+1)
	cache[1] = 1
	for i := 2; i <= n; i++ {
		pos := 2*i - 1
		cache[i] = (sum[pos] * cache[i-1]) % mod1359
	}
	return cache[n]
}

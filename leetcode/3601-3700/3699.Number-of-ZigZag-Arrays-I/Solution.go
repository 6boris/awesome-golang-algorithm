package Solution

func Solution(n int, l int, r int) int {
	const MOD = 1_000_000_007

	dp0 := make([]int, r+1)
	dp1 := make([]int, r+1)
	sum0 := make([]int, r+2)
	sum1 := make([]int, r+2)

	for i := l; i <= r; i++ {
		dp0[i] = 1
		dp1[i] = 1
		val := i - l + 1
		sum0[i] = val
		sum1[i] = val
	}

	for i := 1; i < n; i++ {
		for j := l; j <= r; j++ {
			dp0[j] = (sum1[r] - sum1[j] + MOD) % MOD
			dp1[j] = sum0[j-1]
		}

		sum0[l] = dp0[l]
		sum1[l] = dp1[l]
		for j := l + 1; j <= r; j++ {
			sum0[j] = (sum0[j-1] + dp0[j]) % MOD
			sum1[j] = (sum1[j-1] + dp1[j]) % MOD
		}
	}

	return (sum0[r] + sum1[r]) % MOD
}

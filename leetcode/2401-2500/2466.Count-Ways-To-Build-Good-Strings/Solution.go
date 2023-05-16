package Solution

const mod2466 = 1000000007

func Solution(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	ans := 0
	dp[zero] += 1
	dp[one] += 1
	// 0, 1, 1
	i := zero
	if one < i {
		i = one
	}
	for ; i <= high; i++ {
		if pre := i - zero; pre >= 0 {
			dp[i] = (dp[i] + dp[pre]) % mod2466
		}
		if pre := i - one; pre >= 0 {
			dp[i] = (dp[i] + dp[pre]) % mod2466
		}
	}
	for i := low; i <= high; i++ {
		ans = (ans + dp[i]) % mod2466
	}
	return ans
}

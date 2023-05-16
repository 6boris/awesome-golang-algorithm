package Solution

const mod1416 = 1000000007

func Solution(s string, k int) int {
	length := len(s)
	dp := make([]int, length)
	dp[0] = 1
	for idx := 1; idx < length; idx++ {
		if s[idx] != '0' {
			// add self
			dp[idx] = dp[idx-1] % mod1416
		}
		x := 1
		base := int(s[idx] - '0')
		for p := idx - 1; p >= 0; p-- {
			x *= 10
			if x > k {
				break
			}
			if s[p] == '0' {
				//leading zero
				continue
			}
			tmp := int(s[p]-'0')*x + base
			if tmp > k {
				break
			}
			base = tmp
			add := 1
			if p >= 1 {
				add = dp[p-1]
			}
			dp[idx] += add
			dp[idx] %= mod1416
		}
	}
	return dp[length-1]
}

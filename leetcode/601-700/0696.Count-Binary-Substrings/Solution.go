package Solution

func Solution(s string) int {
	dp := make([]int, len(s))
	dp[0] = 1
	ans, pre := 0, 0
	for idx := 1; idx < len(s); idx++ {
		if s[idx] == s[pre] {
			dp[idx] = idx - pre + 1
			continue
		}
		dp[idx] = 1
		pre = idx
	}
	for idx := 1; idx < len(s); idx++ {
		other := idx - dp[idx]
		if other < 0 || dp[other] < dp[idx] {
			continue
		}
		ans++
	}
	return ans
}

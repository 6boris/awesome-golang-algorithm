package Solution

func Solution(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	dp := make([]int, 26)
	dp[s[0]-'a'] = 1
	c := 1
	for i := 1; i < length; i++ {
		if s[i]-s[i-1] == 1 || s[i-1]-s[i] == 25 {
			//ab, za
			c++
		} else {
			c = 1
		}
		dp[s[i]-'a'] = max(dp[s[i]-'a'], c)
	}
	ans := 0
	for i := 0; i < 26; i++ {
		ans += dp[i]
	}
	return ans
}

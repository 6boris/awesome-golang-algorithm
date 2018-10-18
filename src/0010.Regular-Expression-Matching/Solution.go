package Solution

func isMatch(s string, p string) bool {
	// DP解法
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			// 检查每个匹配的第一个字母是否匹配
			fm := false
			if i < len(s) && (s[i] == p[j] || p[j] == '.') {
				fm = true
			}

			// 当第二个字符模式中有*时
			if (j+1) < len(p) && p[j+1] == '*' {
				// 考虑*为0并且重合
				// 如果不匹配，检查第一个匹配是否与[i + 1，j]匹配
				dp[i][j] = dp[i][j+2] || (fm && dp[i+1][j])
				//fmt.Println("a",i,j,dp[i][j])
			} else {
				dp[i][j] = fm && dp[i+1][j+1]
				//fmt.Println("b",i,j,dp[i][j])
			}
		}
	}
	return dp[0][0]
}

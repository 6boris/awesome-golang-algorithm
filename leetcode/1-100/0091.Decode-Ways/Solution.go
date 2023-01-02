package Solution

func numDecodings(s string) int {
	// 思路：
	// 状态表示：f[i] i个数（最后一个数字是s[i-1]）解码得到的所有字符串的数量 f[i]从0开始
	// 状态计算： 集合可以分成：最后一个i是一位数，和最后一个i是两位数
	// 之前想过的错误思路：f[i] 表示以i结尾的decode共有多少种解法
	// 集合划分为不算i：f[i - j] + f[j] 和算i: f[i];
	n := len(s)
	// 加一是为了方便边界特判
	f := make([]int, n+1)
	// 空字符串特判
	f[0] = 1
	// 如果最后一个i是个位数 1 ~ 26
	// 难点：如何转化char为int
	for i := 1; i <= n; i++ { // 从一开始是因为表示的是前多少个字母，不是下标从一开始的字母
		// 如果最后一个i是个位数，且不为0
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		// 如果最后一个i是两位数，且在1~26之间
		if i >= 2 {
			t := (s[i-2]-'0')*10 + s[i-1] - '0'
			if t >= 10 && t <= 26 {
				f[i] += f[i-2]
			}
		}
	}
	return f[n]
}

func numDecodings1(s string) int {
	l := len(s)
	// 状态只依赖于前一个或者前2个，所以内存可以优化成O(1)
	dp := make([]int, l)
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	for idx := 1; idx < l; idx++ {
		if s[idx] == '0' {
			// 0只可以与前面的一个字符组合成10, 20,
			if s[idx-1] != '1' && s[idx-1] != '2' {
				return 0
			}
			// 10 这种情况
			i := idx - 2
			if i < 0 {
				i = 0
			}
			dp[idx] = dp[i]
			continue
		}

		dp[idx] = dp[idx-1]
		// 108 这种情况就是将8放到最后就可以，
		if s[idx-1] == '0' {
			continue
		}
		// 112, 如果要与前一个字符组成1-26， 就需要s[idx-1] =1 或者2
		// s[idx-1] = 1, s[idx]是什么都可以11-19
		// s[idx-1] = 2, s[idx] 只能是21-26.
		if (s[idx-1] == '1') || s[idx-1] == '2' && s[idx] >= '1' && s[idx] <= '6' {
			i := idx - 2
			if i < 0 {
				i = 0
			}
			dp[idx] += dp[i]
		}
	}

	return dp[l-1]
}

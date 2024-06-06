package Solution

func Solution(n int) int {
	if n == 1 {
		return 0
	}
	// 存在一个当前复制几个问题
	// 1  复制了1次
	// dp[i][j] 表示通过复制j个得到i个元素需要的最少次数，0就是无法到达
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		if i > 1 {
			for j := 0; j <= n; j++ {
				dp[i][j] = -1
			}
		}
	}
	for i := 2; i <= n; i++ {
		// 通过粘贴nci 到达i
		for ct := i - 1; ct > 0; ct-- {
			start := i - ct // 从那个位置开始
			if start > ct && dp[start][ct] != -1 {
				if r := dp[start][ct] + 1; dp[i][ct] == -1 || dp[i][ct] > r {
					dp[i][ct] = r
				}
			}
			if start == ct {
				op := -1
				if start == 1 {
					op = 0
				}
				for pn := 1; pn < start; pn++ {
					if dp[start][pn] != -1 && (op == -1 || dp[start][pn] < op) {
						op = dp[start][pn]
					}
				}
				op += 2 //copy and past
				if dp[i][ct] == -1 || dp[i][ct] > op {
					dp[i][ct] = op
				}
			}
		}
	}
	op := -1
	for i := 1; i < n; i++ {
		if dp[n][i] == -1 {
			continue
		}
		if op == -1 || dp[n][i] < op {
			op = dp[n][i]
		}
	}
	return op
}

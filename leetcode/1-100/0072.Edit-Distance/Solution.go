package Solution

/*
	设状态为dp[i][j],表示A[0,i]与B[0,j]的最小编辑距离，对应形式分别为str1c,str2d
	如果c==d,f[i][j]=f[i-1][j-1]
	如果c!=d,
	a. 如果将c换成d,则f[i][j]=f[i-1][j-1]+1
	b. 如果在c后面加一个d,f[i][j]=f[i][j-1]+1
	c. 如果删掉c,f[i][j]=f[i-1][j]+1
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := [][]int{}
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
		}
	}
	return dp[m][n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

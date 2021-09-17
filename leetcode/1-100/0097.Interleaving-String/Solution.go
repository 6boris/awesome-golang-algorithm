package Solution

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s2)+len(s1) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(s2)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true

	for i := 0; i < len(s1)+1; i++ {
		for j := 0; j < len(s2)+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true
	for i := 1; i < len(s1)+1; i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 1; j < len(s2)+1; j++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			dp[i][j] = (s1[i-1] == s3[i+j-1] && dp[i-1][j]) || (s2[j-1] == s3[i+j-1] && dp[i][j-1])
		}
	}
	return dp[len(s1)][len(s2)]
}

func p(x [][]bool) {
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}

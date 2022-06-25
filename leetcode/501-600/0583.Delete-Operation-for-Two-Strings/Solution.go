package Solution

func Solution(word1, word2 string) int {
	dp := make([][]int, 2)
	for idx := 0; idx < 2; idx++ {
		dp[idx] = make([]int, len(word2)+1)
	}
	for idx := 0; idx < len(word2)+1; idx++ {
		dp[0][idx] = idx
	}
	loop := 1
	for _, _byte := range []byte(word1) {
		dp[loop][0] = dp[1-loop][0] + 1
		for col := 1; col <= len(word2); col++ {
			leftTop := dp[1-loop][col-1]
			left := dp[loop][col-1]
			top := dp[1-loop][col]
			dp[loop][col] = leftTop
			if _byte != word2[col-1] {
				dp[loop][col] += 2
				if left+1 < dp[loop][col] {
					dp[loop][col] = left + 1
				}
				if top+1 < dp[loop][col] {
					dp[loop][col] = top + 1
				}
			}
		}

		loop = 1 - loop
	}
	return dp[1-loop][len(word2)]
}

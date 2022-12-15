package Solution

func Solution(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	dp := make([][]int, 2)
	for idx := 0; idx < 2; idx++ {
		dp[idx] = make([]int, l2+1)
	}
	loop := 1
	for row := 0; row < l1; row++ {
		for col := 1; col <= l2; col++ {
			if text1[row] == text2[col-1] {
				dp[loop][col] = dp[1-loop][col-1] + 1
				continue
			}
			_max := dp[1-loop][col]
			if dp[loop][col-1] > _max {
				_max = dp[loop][col-1]
			}
			dp[loop][col] = _max
		}

		loop = 1 - loop
	}
	return dp[1-loop][l2]
}

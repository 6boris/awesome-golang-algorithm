package Solution

const mod1444 = 1000000007

func Solution(pizza []string, k int) int {
	rows := len(pizza)
	cols := len(pizza[0])
	dp := make([][][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([][]int, rows)
		for r := 0; r < rows; r++ {
			dp[i][r] = make([]int, cols)
		}
	}

	apples := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		apples[i] = make([]int, cols+1)
	}
	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			isApple := 0
			if pizza[row][col] == 'A' {
				isApple++
			}
			apples[row][col] = isApple + apples[row+1][col] + apples[row][col+1] - apples[row+1][col+1]
			if apples[row][col] > 0 {
				dp[0][row][col] = 1
			}
		}
	}

	for remain := 1; remain < k; remain++ {
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				for move := row + 1; move < rows; move++ {
					if r := apples[row][col] - apples[move][col]; r > 0 {
						dp[remain][row][col] += dp[remain-1][move][col]
						dp[remain][row][col] %= mod1444
					}
				}

				for move := col + 1; move < cols; move++ {
					if r := apples[row][col] - apples[row][move]; r > 0 {
						dp[remain][row][col] += dp[remain-1][row][move]
						dp[remain][row][col] %= mod1444
					}
				}
			}
		}
	}
	return dp[k-1][0][0]
}

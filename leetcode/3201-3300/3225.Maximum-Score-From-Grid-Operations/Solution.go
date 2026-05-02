package Solution

func Solution(grid [][]int) int64 {
	n := len(grid[0])
	if n == 1 {
		return 0
	}

	dp := make([][][]int64, n)
	for i := range dp {
		dp[i] = make([][]int64, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, n+1)
		}
	}

	prevMax := make([][]int64, n+1)
	for i := range prevMax {
		prevMax[i] = make([]int64, n+1)
	}
	prevSuffixMax := make([][]int64, n+1)
	for i := range prevSuffixMax {
		prevSuffixMax[i] = make([]int64, n+1)
	}

	colSum := make([][]int64, n)
	for c := 0; c < n; c++ {
		colSum[c] = make([]int64, n+1)
		for r := 1; r <= n; r++ {
			colSum[c][r] = colSum[c][r-1] + int64(grid[r-1][c])
		}
	}

	for i := 1; i < n; i++ {
		for currH := 0; currH <= n; currH++ {
			for prevH := 0; prevH <= n; prevH++ {
				if currH <= prevH {
					extraScore := colSum[i][prevH] - colSum[i][currH]
					dp[i][currH][prevH] = max(dp[i][currH][prevH], prevSuffixMax[prevH][0]+extraScore)
				} else {
					extraScore := colSum[i-1][currH] - colSum[i-1][prevH]
					dp[i][currH][prevH] = max(dp[i][currH][prevH],
						max(prevSuffixMax[prevH][currH], prevMax[prevH][currH]+extraScore))
				}
			}
		}

		for currH := 0; currH <= n; currH++ {
			prevMax[currH][0] = dp[i][currH][0]
			for prevH := 1; prevH <= n; prevH++ {
				var penalty int64 = 0
				if prevH > currH {
					penalty = colSum[i][prevH] - colSum[i][currH]
				}
				prevMax[currH][prevH] = max(prevMax[currH][prevH-1], dp[i][currH][prevH]-penalty)
			}

			prevSuffixMax[currH][n] = dp[i][currH][n]
			for prevH := n - 1; prevH >= 0; prevH-- {
				prevSuffixMax[currH][prevH] = max(prevSuffixMax[currH][prevH+1], dp[i][currH][prevH])
			}
		}
	}

	var ans int64 = 0
	for k := 0; k <= n; k++ {
		ans = max(ans, max(dp[n-1][n][k], dp[n-1][0][k]))
	}

	return ans
}

package Solution

func Solution(stoneValue []int) int {
	n := len(stoneValue)

	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + stoneValue[i]
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == j {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		maxScore := 0
		for k := i; k < j; k++ {
			leftSum := prefixSum[k+1] - prefixSum[i]
			rightSum := prefixSum[j+1] - prefixSum[k+1]

			var currentScore int
			if leftSum < rightSum {
				currentScore = leftSum + dfs(i, k)
			} else if leftSum > rightSum {
				currentScore = rightSum + dfs(k+1, j)
			} else {
				score1 := leftSum + dfs(i, k)
				score2 := rightSum + dfs(k+1, j)
				if score1 > score2 {
					currentScore = score1
				} else {
					currentScore = score2
				}
			}

			if currentScore > maxScore {
				maxScore = currentScore
			}
		}

		memo[i][j] = maxScore
		return maxScore
	}

	return dfs(0, n-1)
}

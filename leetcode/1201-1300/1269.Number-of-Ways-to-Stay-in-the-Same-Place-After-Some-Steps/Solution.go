package Solution

const mod1269 = 1000000007

func Solution(steps int, arrLen int) int {
	dp := make([][]int, arrLen)
	if steps <= arrLen {
		arrLen = steps
	}
	for i := 0; i < arrLen; i++ {
		dp[i] = make([]int, steps+1)
	}
	for i := 0; i < arrLen; i++ {
		for j := 0; j <= steps; j++ {
			dp[i][j] = -1
		}
	}
	var dirs = []int{-1, 0, 1}
	var dfs func(int, int) int
	dfs = func(index, useSteps int) int {
		if useSteps == 0 {
			if index == 0 {
				return 1
			}
			return 0
		}
		if index >= arrLen {
			return 0
		}
		if dp[index][useSteps] != -1 {
			return dp[index][useSteps]
		}
		ans := 0
		for _, dir := range dirs {
			if ni := index + dir; ni >= 0 {
				ans = (ans + dfs(ni, useSteps-1)) % mod1269
			}
		}
		dp[index][useSteps] = ans
		return ans
	}
	return dfs(0, steps)
}

package Solution

import "fmt"

type mm struct {
	max, min float32

	maxStr, minStr string
}

func Solution(nums []int) string {
	if len(nums) == 1 {
		return fmt.Sprintf("%d", nums[0])
	}
	n := len(nums)
	dp := make([][]mm, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]mm, n)
		s := fmt.Sprintf("%d", nums[i])
		dp[i][i] = mm{max: float32(nums[i]), min: float32(nums[i]), maxStr: s, minStr: s}
	}

	for l := 2; l <= n; l++ {
		for end := l - 1; end < n; end++ {
			start := end - l + 1
			for pre := end; pre > start; pre-- {
				m1 := dp[start][pre-1].max
				m2 := dp[pre][end].min
				if r := m1 / m2; r > dp[start][end].max {
					dp[start][end].max = r
					b := dp[pre][end].minStr
					if pre != end {
						b = "(" + b + ")"
					}
					dp[start][end].maxStr = dp[start][pre-1].maxStr + "/" + b
				}
			}

			for pre := end; pre > start; pre-- {
				m1 := dp[start][pre-1].min
				m2 := dp[pre][end].max
				if r := m1 / m2; dp[start][end].min == 0 || r < dp[start][end].min {
					dp[start][end].min = r
					b := dp[pre][end].maxStr
					if pre != end {
						b = "(" + b + ")"
					}
					dp[start][end].minStr = dp[start][pre-1].minStr + "/" + b
				}
			}
		}
	}
	/*
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				fmt.Printf("----- %d - %d, %+v\n", i, j, dp[i][j])
			}
		}
	*/
	return dp[0][n-1].maxStr
}

package Solution

func sort1626(scores, ages []int) {
	l := len(scores)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if (ages[j] > ages[j+1]) || (ages[j] == ages[j+1] && scores[j] > scores[j+1]) {
				ages[j], ages[j+1] = ages[j+1], ages[j]
				scores[j], scores[j+1] = scores[j+1], scores[j]
			}
		}
	}
}
func Solution(scores []int, ages []int) int {
	sort1626(scores, ages)
	l := len(scores)
	dp := make([]int, l)
	ans := 0
	for i := 0; i < l; i++ {
		dp[i] = scores[i]
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	for i := 1; i < l; i++ {
		for pre := i - 1; pre >= 0; pre-- {
			if scores[i] >= scores[pre] {
				if r := dp[pre] + scores[i]; r > dp[i] {
					dp[i] = r
				}
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

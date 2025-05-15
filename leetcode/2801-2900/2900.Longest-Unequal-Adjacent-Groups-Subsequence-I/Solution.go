package Solution

func Solution(words []string, groups []int) []string {
	l, index := 1, 0
	dp := make([][2]int, len(words))
	dp[0] = [2]int{1, -1} // 前一个
	for i := 1; i < len(words); i++ {
		dp[i] = [2]int{1, -1}
		for pre := i - 1; pre >= 0; pre-- {
			if groups[i] != groups[pre] {
				if dp[pre][0]+1 > dp[i][0] {
					dp[i] = [2]int{dp[pre][0] + 1, pre}
				}
			}
		}
		if dp[i][0] > l {
			l = dp[i][0]
			index = i
		}
	}
	ans := make([]string, l)
	for i := l - 1; i >= 0; i-- {
		ans[i] = words[index]
		index = dp[index][1]
	}
	return ans
}

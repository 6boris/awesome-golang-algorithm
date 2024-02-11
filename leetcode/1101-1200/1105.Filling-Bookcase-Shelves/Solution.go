package Solution

func Solution(books [][]int, shelfWidth int) int {
	dp := make([]int, len(books))
	dp[0] = books[0][1]
	for i := 1; i < len(books); i++ {
		dp[i] = books[i][1] + dp[i-1]
		tmpMax := books[i][1]
		curWidth := books[i][0]
		for pre := i - 1; pre >= 0; pre-- {
			curWidth += books[pre][0]
			if curWidth > shelfWidth {
				break
			}
			if books[pre][1] > tmpMax {
				tmpMax = books[pre][1]
			}
			add := 0
			if pre > 0 {
				add = dp[pre-1]
			}
			if r := add + tmpMax; r < dp[i] {
				dp[i] = r
			}
		}
	}
	return dp[len(books)-1]
}

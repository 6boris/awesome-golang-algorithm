package Solution

func Solution(piles [][]int, k int) int {
	l := len(piles)
	dp := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i <= l; i++ {
		for c := 0; c <= k; c++ {
			sum := 0
			for cur := 0; cur <= c && cur <= len(piles[i-1]); cur++ {
				if cur > 0 {
					sum += piles[i-1][cur-1]
				}
				if r := dp[i-1][c-cur] + sum; r > dp[i][c] {
					dp[i][c] = r
				}
			}
		}
	}
	return dp[l][k]
}

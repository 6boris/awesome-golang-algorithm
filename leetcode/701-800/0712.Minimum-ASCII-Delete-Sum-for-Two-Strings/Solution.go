package Solution

func Solution(s1 string, s2 string) int {
	l1, l2 := len(s1), len(s2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {

		dp[i] = make([]int, l2+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
		for j := 1; j <= l2; j++ {
			dp[i][j] = dp[i][j-1] + int(s2[j-1])
		}
	}
	for r := 1; r <= l1; r++ {
		for c := 1; c <= l2; c++ {
			x := dp[r-1][c-1]
			if s1[r-1] == s2[c-1] {
				dp[r][c] = x
				continue
			}
			y := dp[r-1][c]
			z := dp[r][c-1]
			a := x + int(s1[r-1]) + int(s2[c-1])
			if b := y + int(s1[r-1]); r < a {
				a = b
			}
			if b := z + int(s2[c-1]); b < a {
				a = b
			}
			dp[r][c] = a
		}
	}
	return dp[l1][l2]
}

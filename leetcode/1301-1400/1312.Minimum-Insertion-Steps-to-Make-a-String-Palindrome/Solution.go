package Solution

func Solution(s string) int {
	length := len(s)
	one := []byte(s)
	for s, e := 0, length-1; s < e; s, e = s+1, e-1 {
		one[s], one[e] = one[e], one[s]
	}
	dp := make([][]int, length+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([]int, length+1)
	}
	for row := 1; row <= length; row++ {
		for col := 1; col <= length; col++ {
			if one[col-1] == s[row-1] {
				dp[row][col] = dp[row-1][col-1] + 1
			} else {
				x := dp[row-1][col]
				if y := dp[row][col-1]; y > x {
					x = y
				}
				dp[row][col] = x
			}
		}
	}
	return length - dp[length][length]
}

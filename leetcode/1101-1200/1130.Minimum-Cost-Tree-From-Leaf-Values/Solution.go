package Solution

type node1130 struct {
	maximum, sum int
}

func Solution(arr []int) int {
	la := len(arr)
	dp := make([][]node1130, la)
	for i := 0; i < la; i++ {
		dp[i] = make([]node1130, la)
		dp[i][i] = node1130{maximum: arr[i]}
	}

	for l := 2; l <= la; l++ {
		// 0, 1, 2
		for end := l - 1; end < la; end++ {
			start := end + 1 - l
			for k := start; k < end; k++ {
				nodeValue := dp[start][k].maximum * dp[k+1][end].maximum
				sum := dp[start][k].sum + dp[k+1][end].sum + nodeValue
				m := dp[start][k].maximum
				if dp[k+1][end].maximum > m {
					m = dp[k+1][end].maximum
				}
				if dp[start][end].sum == 0 || dp[start][end].sum > sum {
					dp[start][end].maximum = m
					dp[start][end].sum = sum
				}
			}
		}
	}
	return dp[0][la-1].sum
}

package Solution

func Solution(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
    dp := make([][]int, n + 1)
    for i, _ := range dp {
        dp[i] = make([]int, m + 1)
    }
    res := 0
    for i := 1; i < n + 1; i++ {
        for j := 1; j < m + 1; j++ {
            if matrix[i - 1][j - 1] == 1 {
                dp[i][j] = 1 + min(dp[i - 1][j - 1], min(dp[i - 1][j], dp[i][j - 1]))
            }
            res += dp[i][j]
        }
    }
    return res
}

func min(a, b int) int {
    if a < b { return a }
    return b
}

package Solution

import "slices"

func Solution(points [][]int) int64 {
	m, n := len(points), len(points[0])

	dp := make([]int64, n)
	for i := 0; i < n; i++ {
		//  第一行的初始化
		dp[i] = int64(points[0][i])
	}
	// 如果按照dp思路去做，我们就需要从第二行开始，遍历到最后，
	// 对于每一行的每一个元素，都需要与上一行去做计算，得到points[i][j]的最大分数, 这样复杂度直接到n^3了。
	// 如果当前在计算第二行第四列的元素,
	// dp[2][3] = max(dp[1][0]-3, dp[1][1]-2, dp[1][2]-1, dp[1][3])
	// dp[2][1] = max(dp[1][0]-2, dp[1][1]-1, dp[1][2])
	// dp[2][1] = max(dp[1][0]-1, dp[1][1])
	// 有大量的的计算,因为有些数据再减去分数后，会越来越小，计算就没有意义了。
	// 对于每一行, 我们用left和right两个数组分别计算出[0:i]的最大值和[i:length]的最大值，最后二者取最大，+points就是最大值
	for i := 1; i < m; i++ {
		next := make([]int64, n)
		left, right := make([]int64, n), make([]int64, n)
		left[0] = dp[0]
		for c := 1; c < n; c++ {
			left[c] = max(left[c-1]-1, dp[c])
		}
		right[n-1] = dp[n-1]
		for c := n - 2; c >= 0; c-- {
			right[c] = max(right[c+1]-1, dp[c])
		}
		for c := 0; c < n; c++ {
			next[c] = max(left[c], right[c]) + int64(points[i][c])
		}
		copy(dp, next)
	}

	return slices.Max(dp)
}

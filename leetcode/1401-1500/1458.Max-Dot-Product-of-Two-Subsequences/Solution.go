package Solution

import (
	"math"
)

func Solution(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	dp := make([][]int, l1)
	for i := 0; i < l1; i++ {
		dp[i] = make([]int, l2)
		for j := 0; j < l2; j++ {
			dp[i][j] = math.MinInt
		}
	}
	// -1, -1
	// 1, 1
	//bool是否确定返回的0是否是因为到了边界导致的, 上面的测试用例就是例子，应该返回-1，但是遇到了边界导致返回0，结果不对
	var dfs func(int, int) (int, bool)
	dfs = func(i, j int) (int, bool) {
		if i == l1 || j == l2 {
			return 0, false
		}
		if dp[i][j] != math.MinInt {
			return dp[i][j], true
		}
		a := nums1[i] * nums2[j]

		if b, ok := dfs(i+1, j+1); ok && b > 0 {
			a += b
		}
		// 因为到了边界，返回的0，导致比-1大
		if c, ok := dfs(i, j+1); ok && c > a {
			a = c
		}
		if d, ok := dfs(i+1, j); ok && d > a {
			a = d
		}
		dp[i][j] = a
		return a, true
	}

	a, _ := dfs(0, 0)
	return a
}

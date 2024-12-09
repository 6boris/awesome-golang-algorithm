package Solution

func Solution(nums []int, queries [][]int) []bool {
	lq := len(queries)
	ans := make([]bool, lq)
	l := len(nums)
	dp := make([]int, l)
	for i := 1; i < l; i++ {
		cur := nums[i] & 1
		pre := nums[i-1] & 1
		if cur != pre {
			dp[i] = dp[i-1] + 1
		}
	}
	// 3, 4, 1, 2, 6
	// 0, 1, 2, 3, 0
	for i, q := range queries {
		s, e := q[0], q[1]
		if s == e || s >= e-dp[e] {
			ans[i] = true
		}
	}
	return ans
}

package Solution

func Solution(arr []int, k int) int {
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		m := 0
		for pre := i; pre >= 0 && pre > i-k; pre-- {
			if arr[pre] > m {
				m = arr[pre]
			}
			l := i - pre + 1
			sum := m * l
			add := 0
			if pre > 1 {
				add = dp[pre-1]
			}
			if r := sum + add; r > dp[i] {
				dp[i] = r
			}
		}
	}
	return dp[len(arr)-1]
}

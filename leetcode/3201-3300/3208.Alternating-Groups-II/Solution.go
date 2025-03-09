package Solution

func Solution(colors []int, k int) int {
	l := len(colors)
	dp := make([]int, len(colors))
	dp[0] = 1
	for i := 1; i < l; i++ {
		if colors[i] != colors[i-1] {
			dp[i] = dp[i-1] + 1
			continue
		}
		dp[i] = 1
	}
	ans := 0
	for start := 0; start <= l-k; start++ {
		end := start + k - 1
		if dp[end] >= k {
			ans++
		}
	}
	if colors[0] != colors[l-1] {
		for start := l - k + 1; start < l; start++ {
			end := (start+k)%l - 1
			if dp[l-1] >= l-start && dp[end] >= end+1 {
				ans++
			}
		}
	}

	return ans
}

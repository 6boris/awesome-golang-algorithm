package Solution

const mod923 = 1000000007

func Solution(arr []int, target int) int {
	dp := make([][101]int, len(arr))
	dp[0][arr[0]] = 1
	for i := 1; i < len(arr); i++ {
		for j := 0; j < 101; j++ {
			dp[i][j] = dp[i-1][j]
		}
		dp[i][arr[i]]++
	}

	ans := 0
	for i := 2; i < len(arr); i++ {
		cur := arr[i]
		for pre := i - 1; pre > 0; pre-- {
			if diff := target - arr[pre] - cur; diff >= 0 && diff <= 100 {
				ans = (ans + dp[pre-1][diff]) % mod923
			}
		}
	}
	return ans
}

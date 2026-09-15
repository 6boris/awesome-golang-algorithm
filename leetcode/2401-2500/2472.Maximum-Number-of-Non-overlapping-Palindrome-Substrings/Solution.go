package Solution

func Solution(s string, k int) int {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	for length := 1; length <= n; length++ {
		for left := 0; left+length <= n; left++ {
			right := left + length - 1
			isPalindrome[left][right] = s[left] == s[right] &&
				(length <= 2 || isPalindrome[left+1][right-1])
		}
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		for j := 0; j+k <= i; j++ {
			if isPalindrome[j][i-1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n]
}

package Solution

func Solution(s string) int {
	ans := 0
	dp := make([][]int, len(s))
	for row := 0; row < len(s); row++ {
		dp[row] = make([]int, len(s))
	}

	for idx := len(s) - 1; idx >= 0; idx-- {
		for in := idx; in < len(s); in++ {
			if idx == in || (s[idx] == s[in] && (idx+1 == in || dp[idx+1][in-1] == 1)) {
				dp[idx][in] = 1
				ans++
			}
		}
	}

	return ans
}

func Solution1(s string) int {
	baseLen := len(s)
	ans := baseLen

	sBytes := []byte(s)
	for length := 2; length <= baseLen; length++ {
		for start := 0; start <= baseLen-length; start++ {
			if isPalindromic(sBytes, start, start+length-1) {
				ans++
			}
		}
	}
	return ans
}

func isPalindromic(sBytes []byte, start, end int) bool {
	for s, e := start, end; s < e; s, e = s+1, e-1 {
		if sBytes[s] != sBytes[e] {
			return false
		}
	}
	return true
}

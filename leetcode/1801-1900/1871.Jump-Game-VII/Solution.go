package Solution

func Solution(s string, minJump, maxJump int) bool {
	size := len(s)
	if s[size-1] != '0' {
		return false
	}

	dp := make([]bool, size)
	dp[0] = true
	ps := 0
	for i := 1; i < size; i++ {
		if i >= minJump {
			if dp[i-minJump] {
				ps++
			}
		}
		if i > maxJump {
			if dp[i-maxJump-1] {
				ps--
			}
		}
		if s[i] == '0' && ps > 0 {
			dp[i] = true
		}

	}

	return dp[size-1]
}

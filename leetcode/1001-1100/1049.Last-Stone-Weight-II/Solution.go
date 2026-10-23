package Solution

func Solution(stones []int) int {
	sum := 0
	for _, n := range stones {
		sum += n
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, s := range stones {
		for j := target; j >= s; j-- {
			dp[j] = dp[j] || dp[j-s]
		}
	}

	for j := target; j >= 0; j-- {
		if dp[j] {
			return sum - 2*j
		}
	}
	return 0
}

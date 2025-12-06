package Solution

func Solution(nums []int, k int) int {
	n := len(nums)
	mod := int64(1e9 + 7)
	dp := make([]int64, n+1)
	prefix := make([]int64, n+1)
	minQ := make([]int, 0)
	maxQ := make([]int, 0)

	dp[0] = 1
	prefix[0] = 1

	for i, j := 0, 0; i < n; i++ {
		// maintain the maximum value queue
		for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] <= nums[i] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, i)
		// maintain the minimum value queue
		for len(minQ) > 0 && nums[minQ[len(minQ)-1]] >= nums[i] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, i)

		// adjust window
		for len(maxQ) > 0 && len(minQ) > 0 &&
			nums[maxQ[0]]-nums[minQ[0]] > k {
			if maxQ[0] == j {
				maxQ = maxQ[1:]
			}
			if minQ[0] == j {
				minQ = minQ[1:]
			}
			j++
		}

		if j > 0 {
			dp[i+1] = (prefix[i] - prefix[j-1] + mod) % mod
		} else {
			dp[i+1] = prefix[i] % mod
		}
		prefix[i+1] = (prefix[i] + dp[i+1]) % mod
	}

	return int(dp[n])
}

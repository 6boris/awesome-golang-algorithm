package Solution

func Solution(nums []int) int {
	dp := make([]int, len(nums))
	maxCount := make([]int, len(nums))
	dp[0], maxCount[0] = 1, 1
	longestSubSeq := 1
	for idx := 1; idx < len(nums); idx++ {
		dp[idx] = 1
		maxCount[idx] = 1
		for pre := idx - 1; pre >= 0; pre-- {
			if nums[pre] < nums[idx] {
				times := dp[pre] + 1
				if times == dp[idx] {
					maxCount[idx] += maxCount[pre]
					continue
				}
				if times > dp[idx] {
					dp[idx] = times
					maxCount[idx] = maxCount[pre]
				}
				if longestSubSeq < dp[idx] {
					longestSubSeq = dp[idx]
				}
			}
		}
	}
	result := 0
	for idx := 0; idx < len(nums); idx++ {
		if longestSubSeq == dp[idx] {
			result += maxCount[idx]
		}
	}
	return result
}

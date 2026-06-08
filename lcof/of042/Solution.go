package Solution

// 动态规划1
func maxSubArray_1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

// 动态规划2
func maxSubArray_2(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i-1]+nums[i], nums[i])
		ans = max(ans, nums[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

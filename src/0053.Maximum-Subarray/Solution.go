package Solution

func maxSubArray(nums []int) int {
	dp, ans := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = Max(nums[i], nums[i]+dp[i-1])
		ans = Max(ans, dp[i])
	}

	return ans
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

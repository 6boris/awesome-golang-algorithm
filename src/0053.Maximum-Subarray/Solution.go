package Solution

func maxSubArray(nums []int) int {
	dp := nums[0]
	max := dp

	for i := 1; i < len(nums); i++ {
		if dp > 0 {
			dp = nums[i] + dp
		} else {
			dp = nums[i]
		}
		if dp > max {
			max = dp
		}
	}

	return max
}

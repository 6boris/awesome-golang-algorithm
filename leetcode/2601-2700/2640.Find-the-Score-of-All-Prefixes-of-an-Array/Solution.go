package Solution

func Solution(nums []int) []int64 {
	ans := make([]int64, len(nums))
	m := int64(nums[0])
	ans[0] = int64(nums[0]) + m
	for i := 1; i < len(nums); i++ {
		m = max(m, int64(nums[i]))
		ans[i] = int64(nums[i]) + m + ans[i-1]
	}
	return ans
}

package Solution

func Solution(nums []int) []int {
	l := len(nums)
	sum := make([]int, l)
	sum[0] = nums[0]
	for i := 1; i < l; i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	result := make([]int, len(nums))
	for i := 0; i < l; i++ {
		left, right := 0, 0
		if i > 0 {
			s := sum[i-1]
			left = i*nums[i] - s
		}
		if i < l-1 {
			s := sum[l-1] - sum[i]
			right = s - (l-1-i)*nums[i]
		}
		result[i] = left + right
	}
	return result
}

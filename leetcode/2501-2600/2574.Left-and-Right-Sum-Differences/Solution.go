package Solution

func Solution(nums []int) []int {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = 0, 0
	for i := 0; i < n-1; i++ {
		left[i+1] = nums[i] + left[i]
	}
	for i := n - 1; i > 0; i-- {
		right[i-1] = right[i] + nums[i]
	}
	for i := 0; i < n; i++ {
		left[i] -= right[i]
		if left[i] < 0 {
			left[i] = -left[i]
		}
	}
	return left
}

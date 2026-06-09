package Solution

func Solution(nums []int, k int) int64 {
	a, b := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		a = max(a, nums[i])
		b = min(b, nums[i])
	}

	return int64(k) * int64(a-b)
}

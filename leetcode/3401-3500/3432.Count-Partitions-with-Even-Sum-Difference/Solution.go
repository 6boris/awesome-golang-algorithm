package Solution

func Solution(nums []int) int {
	var ret, diff int
	l := len(nums)
	for i := 1; i < l; i++ {
		nums[i] += nums[i-1]
	}
	// 1, 2, 3, 4
	for i := 0; i < l-1; i++ {
		if diff = nums[l-1] - nums[i] - nums[i]; diff&1 == 0 {
			ret++
		}
	}
	return ret
}

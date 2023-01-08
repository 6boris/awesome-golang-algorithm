package Solution

// 看解析可以用单调栈完成
func Solution(nums []int) int64 {
	length := len(nums)
	max, min := make([]int, length), make([]int, length)
	max[0], min[0] = nums[0], nums[0]
	for i := 1; i < length; i++ {
		max[i] = nums[i-1]
		min[i] = nums[i-1]
		if nums[i] > max[i] {
			max[i] = nums[i]
		}
		if nums[i] < min[i] {
			min[i] = nums[i]
		}
	}

	ans := int64(0)
	for l := 2; l <= length; l++ {
		m, n := max[l-1], min[l-1]
		ans += int64(m - n)
		// start
		for i := 1; i <= length-l; i++ {
			newItem := nums[i+l-1]
			if newItem > m {
				m = newItem
			}
			if newItem < n {
				n = newItem
			}
			ans += int64(m - n)
		}
	}
	return ans
}

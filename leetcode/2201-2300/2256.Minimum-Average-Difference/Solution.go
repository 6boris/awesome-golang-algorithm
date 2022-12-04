package Solution

func Solution(nums []int) int {
	minIndex := -1
	minValue := 0
	length := len(nums)
	if length == 1 {
		return 0
	}
	for idx := 1; idx < length; idx++ {
		nums[idx] += nums[idx-1]
	}

	for idx := 0; idx < length; idx++ {
		left := nums[idx] / (idx + 1)
		right := 0
		if idx != length-1 {
			right = (nums[length-1] - nums[idx]) / (length - idx - 1)
		}
		diff := left - right
		if diff < 0 {
			diff = ^diff + 1
		}
		if minIndex == -1 || minValue > diff {
			minIndex = idx
			minValue = diff
		}
	}
	return minIndex
}

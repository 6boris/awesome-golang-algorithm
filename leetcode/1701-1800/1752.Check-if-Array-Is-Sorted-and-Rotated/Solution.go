package Solution

func Solution(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	idx := 1
	for ; idx < length; idx++ {
		if nums[idx] < nums[idx-1] {
			break
		}
	}

	for i := idx; i < length; i++ {
		if idx == i && nums[i] > nums[0] {
			return false
		}
		if i != idx && !(nums[i] >= nums[i-1] && nums[i] <= nums[0]) {
			return false
		}
	}

	return true
}

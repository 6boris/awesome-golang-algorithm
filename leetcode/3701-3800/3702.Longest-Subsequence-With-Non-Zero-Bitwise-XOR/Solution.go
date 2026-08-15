package Solution

func Solution(nums []int) int {
	xor, size := 0, len(nums)
	allZero := true
	for i := range nums {
		xor ^= nums[i]
		if nums[i] > 0 {
			allZero = false
		}
	}
	if xor > 0 {
		return size
	}
	if allZero {
		return 0
	}
	return size - 1
}

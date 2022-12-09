package Solution

func Solution(nums []int) int {
	hills := 0
	idx := 1
	which := 0
	left := false

	for ; idx < len(nums); idx++ {
		if nums[idx] == nums[idx-1] {
			continue
		}

		if which == 0 {
			left = nums[idx] > nums[idx-1]
			which = 1
			continue
		}
		rigth := nums[idx-1] > nums[idx]
		if (left && rigth) || (!left && !rigth) {
			hills++
		}
		left = !rigth
	}
	return hills
}

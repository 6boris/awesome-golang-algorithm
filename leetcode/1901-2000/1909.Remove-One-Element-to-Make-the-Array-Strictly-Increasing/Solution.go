package Solution

func Solution(nums []int) bool {
	pivot := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			if pivot == -1 {
				pivot = i
				continue
			}
			return false
		}
	}
	if pivot == -1 {
		return true
	}

	// remove pivot-1
	if pivot-2 >= 0 && nums[pivot] > nums[pivot-2] {
		return true
	}
	// remove pivot
	if pivot < len(nums)-1 && nums[pivot+1] > nums[pivot-1] {
		return true
	}

	return pivot == len(nums)-1 || pivot == 1
}

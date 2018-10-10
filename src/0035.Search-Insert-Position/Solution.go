package Solution

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := (right + left) >> 1
	for left <= right {
		if target <= nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (right + left) >> 1

	}
	return left
}

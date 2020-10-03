package Solution

func findMin(nums []int) int {
	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		//	数组中最小的元素会小于第一个元素
		if nums[mid] >= nums[0] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}

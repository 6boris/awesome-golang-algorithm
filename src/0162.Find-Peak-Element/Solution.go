package Solution

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) / 2
		if nums[mid] > nums[mid-1] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

package Solution

//	直接扫描
func searchRange(nums []int, target int) []int {
	targetRange := []int{-1, -1}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			targetRange[0] = i
			break
		}
	}
	if targetRange[0] == -1 {
		return targetRange
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] == target {
			targetRange[1] = j
			break
		}
	}
	return targetRange
}

func searchRange2(nums []int, target int) []int {
	targetRange := []int{-1, -1}
	leftIndex := extremeInsertionIndex(nums, target, true)

	if leftIndex == len(nums) || nums[leftIndex] != target {
		return targetRange
	}

	targetRange[0] = leftIndex
	targetRange[1] = extremeInsertionIndex(nums, target, false) - 1

	return targetRange

}
func extremeInsertionIndex(nums []int, target int, left bool) int {
	lo := 0
	hi := len(nums)
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > target || left && target == nums[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

package Solution

func Solution(nums []int) int {
	left, right := nums[0], 1000000000
	for left < right {
		mid := left + (right-left)/2
		buf := int64(0)
		i := 0
		for ; i < len(nums); i++ {
			buf += int64(mid - nums[i])
			if buf < 0 {
				break
			}
		}
		if i == len(nums) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

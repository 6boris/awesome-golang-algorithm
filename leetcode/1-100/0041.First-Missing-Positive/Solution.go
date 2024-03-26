package Solution

func Solution(nums []int) int {
	n := len(nums)
	j := -1
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	j++
	for i := 0; i < j; i++ {
		source := nums[i]
		if source < 0 {
			source = -source
		}
		if source <= j && nums[source-1] > 0 {
			nums[source-1] = -nums[source-1]
		}
	}
	for i := 0; i < j; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return j + 1
}

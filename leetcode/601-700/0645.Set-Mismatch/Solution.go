package Solution

func Solution(nums []int) []int {
	dup := -1
	for _, n := range nums {
		abs := n
		if abs < 0 {
			abs = -abs
		}
		if nums[abs-1] < 0 {
			dup = abs
			continue
		}
		nums[abs-1] *= -1
	}
	missing := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			missing = i + 1
			break
		}
	}
	return []int{dup, missing}
}

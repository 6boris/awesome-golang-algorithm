package Solution

func Solution(nums []int) []int {
	for _, n := range nums {
		cur := n
		if cur < 0 {
			cur = -cur
		}
		if nums[cur-1] > 0 {
			nums[cur-1] = -nums[cur-1]
		}
	}
	missing := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			missing = append(missing, i+1)
		}
	}
	return missing
}

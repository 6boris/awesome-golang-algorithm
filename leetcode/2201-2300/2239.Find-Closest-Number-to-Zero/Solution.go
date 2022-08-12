package Solution

func Solution(nums []int) int {
	absMin, ans := -1, 0
	for idx := 0; idx < len(nums); idx++ {
		target := nums[idx]
		if target < 0 {
			target = -target
		}
		if absMin == -1 {
			// first
			absMin = target
			ans = nums[idx]
			continue
		}
		if target <= absMin {
			if target < absMin || ans < nums[idx] {
				ans = nums[idx]
			}
			absMin = target
		}
	}
	return ans
}

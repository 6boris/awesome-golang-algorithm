package Solution

func Solution(nums []int, n int) int {
	patches := 0
	maxReach := 0
	i := 0

	for maxReach < n {
		if i < len(nums) && nums[i] <= maxReach+1 {
			maxReach += nums[i]
			i++
		} else {
			maxReach = 2*maxReach + 1
			patches++
		}
	}

	return patches
}

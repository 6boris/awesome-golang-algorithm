package Solution

func Solution(nums []int) int {
	minNumbers := make([]int, len(nums))
	minNumbers[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		minNumbers[i] = min(minNumbers[i-1], nums[i])
	}
	ans := -1
	var diff, a int
	for i := len(nums) - 1; i > 0; i-- {
		a = minNumbers[i-1]
		if a >= nums[i] {
			continue
		}
		diff = nums[i] - a
		if ans == -1 || ans < diff {
			ans = diff
		}
	}
	return ans
}

package Solution

func Solution(nums []int) int {
	l := len(nums)
	for i := 1; i < 3; i++ {
		for j := 1; j < l-i; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums[0] + nums[l-1] + nums[l-2]
}

package Solution

func Solution(nums []int) int {
	max, c := 0, 0
	for i := 0; i < len(nums); i++ {
		c = nums[i]*c + nums[i]
		if c > max {
			max = c
		}
	}
	return max
}

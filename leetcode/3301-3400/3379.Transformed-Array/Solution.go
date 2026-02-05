package Solution

func Solution(nums []int) []int {
	l := len(nums)
	ret := make([]int, l)
	steps := 0
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		if nums[i] > 0 {
			steps = nums[i] % l
			ret[i] = nums[(i+steps)%l]
			continue
		}
		steps = (-nums[i]) % l
		ret[i] = nums[(i-steps+l)%l]
	}
	return ret
}

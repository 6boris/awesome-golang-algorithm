package Solution

func Solution(nums []int) int {
	l := len(nums)

	leftMax := make([]int, l)
	rightMin := make([]int, l)
	leftMax[0] = nums[0]
	for i := 1; i < l; i++ {
		leftMax[i] = max(nums[i-1], leftMax[i-1])
	}
	rightMin[l-1] = nums[l-1]
	for i := l - 2; i >= 0; i-- {
		rightMin[i] = min(rightMin[i+1], nums[i+1])
	}
	ret := 0
	for i := 1; i < l-1; i++ {
		lm := leftMax[i]
		rm := rightMin[i]
		if nums[i] > lm && nums[i] < rm {
			ret += 2
			continue
		}
		if nums[i] > nums[i-1] && nums[i] < nums[i+1] {
			ret++
		}
	}
	return ret
}

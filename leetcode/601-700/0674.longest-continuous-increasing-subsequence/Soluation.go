package Soluation

func findLengthOfLCIS_1(nums []int) int {
	ans, start := 0, 0
	for i, v := range nums {
		if i > 0 && v <= nums[i-1] {
			start = i
		}
		ans = max(ans, i-start+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
